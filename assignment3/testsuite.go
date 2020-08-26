package main

//export PATH=$PATH:/usr/local/go/bin

import (
	"fmt" //block.go
	"encoding/hex" //block.go
	"crypto/sha256" //block.go
	"bytes" //blockchain.go
)

/////////////////
//work_queue.go//
/////////////////

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{})
	for i := uint(0); i < nWorkers; i++ {
		go func(){
			q.worker()
		}()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	for i := range queue.Jobs {
		task := i
		result := task.Run()
		queue.Results <- result
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
}

////////////
//block.go//
////////////

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	block := Block{}
	block.PrevHash = []byte("\x00")
	block.Generation = 0
	block.Difficulty = difficulty
	block.Data = ""
	return block 
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	block := Block{}
	block.PrevHash = prev_block.Hash
	block.Generation = prev_block.Generation + 1
	block.Difficulty = prev_block.Difficulty
	block.Data = data
	return block 
}

// Calculate the block's hash.
//support functions
 //fillHexString: fill string with 00 representing each byte in hex
func fillHexString(hexString string, str_len int) string {
	for len(hexString) < str_len {
		hexString = "00" + hexString
	}
	return hexString
}
func (blk Block) CalcHash() []byte {
	//prepend 00 to each empty byte in PrevHash
	prevHash_string := fillHexString(hex.EncodeToString(blk.PrevHash),64)
	block_string := prevHash_string + ":" + fmt.Sprintf("%d",blk.Generation) + ":" + fmt.Sprintf("%d",blk.Difficulty) + ":" + blk.Data + ":" + fmt.Sprintf("%d",blk.Proof)
	hash := sha256.New()
	hash.Write([]byte(block_string))
	hashVal := hash.Sum(nil)
	return hashVal
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	hashLen := len(blk.Hash)
	if(hashLen < int(blk.Difficulty)){return false}
	nBytes := int(blk.Difficulty/8)
	nBits := int(blk.Difficulty%8)
	for i := 0; i < nBytes; i++ {
		if(blk.Hash[hashLen-1-i] != '\x00'){return false}
	}
	bitByte := int(blk.Hash[hashLen-1-nBytes])
	for i := 0; i < nBits; i++ {
		if((bitByte >> i)%2 != 0){return false}
	}
	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}

/////////////
//mining.go//
/////////////

type miningWorker struct {
	RangeStart uint64
	RangeEnd uint64
 	Block Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

func newWorker(start uint64, end uint64, blk Block) miningWorker {
	w := new(miningWorker)
	w.RangeStart = start
	w.RangeEnd = end
	w.Block = blk
	return *w
}

func (w miningWorker) Run() interface{} {
	blk := w.Block
	result := MiningResult{w.RangeStart,false}
	for i:= w.RangeStart; i <= w.RangeEnd; i++ {
		blk.SetProof(i)
		if(blk.ValidHash()) {
			result.Proof = i
			result.Found = true
			return result
		}
	}
	return result
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	//initialize workQueue
	q := Create(uint(workers), uint(chunks))
	chunk_size := uint64((end-start+1)/chunks)
	if(chunk_size < 1){chunk_size = 1};
	//populate workQueue.Jobs with tasks
	for i := uint64(0); i < chunks; i++ {
		chunk_start := start + i*chunk_size
		chunk_end := chunk_start + chunk_size
		if(chunk_end > end){chunk_end = end+1} //ensure MineRange is not exceeded
		q.Enqueue(newWorker(chunk_start,chunk_end,blk))
	}
	//look for valid proof
	result := MiningResult{start,false}
	nResults := uint(0)
	for r := range (q.Results) {
		nResults += 1
		res := r.(MiningResult)
		if res.Found {
			q.Shutdown()
			return res
		}
	}
	return result
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

/////////////////
//blockchain.go//
/////////////////

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain,blk)
}

func (chain Blockchain) IsValid() bool {
	validity := true
	chainyChain := chain.Chain
	//check for empty chain
	if(len(chainyChain) < 1){panic("empty chain!")}

	//check The initial block has previous hash all null bytes and is generation zero.
	initBlock := chainyChain[0]
	if(len(chainyChain[0].PrevHash) < 1) {
		return false
	}
	if(initBlock.PrevHash[0] != 0){validity = false}
	//check The initial block has valid and correct hash value
	if(bytes.Compare(initBlock.Hash, initBlock.CalcHash()) != 0){validity = false}
	if(!initBlock.ValidHash()){validity = false}

	prevBlock := initBlock
	for i := 1; i < len(chainyChain); i++ {
		currBlock := chainyChain[i]
		//check Each block has the same difficulty value.
		if(currBlock.Difficulty != initBlock.Difficulty){validity = false; break}
		//check Each block has a generation value that is one more than the previous block.
		if(currBlock.Generation != (prevBlock.Generation+1)){validity = false; break}
		//check Each block's previous hash matches the previous block's hash.
		if(bytes.Compare(currBlock.PrevHash, prevBlock.Hash) != 0){validity = false; break}
		//check Each block's hash value actually matches its contents.
		if(bytes.Compare(currBlock.Hash, currBlock.CalcHash()) != 0){validity = false; break}
		//check Each block's hash value ends in difficulty null bits.
		if(!currBlock.ValidHash()){validity = false; break}
		prevBlock = currBlock
	}
	return validity
}

func main() {
	/*
	//block.go tests
	b0 := Initial(16)
	//Block Initial() test
	fmt.Println(b0.PrevHash)
	fmt.Println(b0.Generation)
	fmt.Println(b0.Difficulty)
	fmt.Println(b0.Data)
	//Block Next() test
	b1 := b0.Next("message")
	fmt.Println(b1.PrevHash)
	fmt.Println(b1.Generation)
	fmt.Println(b1.Difficulty)
	fmt.Println(b1.Data)
	fmt.Println("\n")
	//Block CalcHash() test
	b0.SetProof(56231)
	fmt.Printf("b0 hash is: %x\n", b0.Hash)
	b1 = b0.Next("message")
	b1.SetProof(2159)
	fmt.Printf("b1 hash is: %x\n", b1.Hash)
	fmt.Println("\n")
	//Block ValidHash() test
	b3 := Initial(16) //Hash = 0
	if(!b3.ValidHash()){fmt.Println("ValidHash test1 passed")}
	b3.SetProof(56231) //Hash = 6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000
	if(b3.ValidHash()){fmt.Println("ValidHash test2 passed")}
	b3.Difficulty = 16 //nBytes = 2, nBits = 0 evaluate nBytes only and passes 0x0000
	if(b3.ValidHash()){fmt.Println("ValidHash test3 passed")}
	b3.Difficulty = 17 //nBytes = 2, nBits = 1, fails evaluating nBits 0xbd = 10111101
	if(!b3.ValidHash()){fmt.Println("ValidHash test4 passed")}
	b3.Hash = b1.Hash //Hash = 9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000
	b3.Difficulty = 18 //nBytes = 2, nBits = 2, passes evaluating nBits 0xf4 = 11110100
	if(b3.ValidHash()){fmt.Println("ValidHash test5 passed")}
	b3.Difficulty = 19 //nBytes = 2, nBits = 3, fails evaluating nBits 0xf4 = 11110100
	if(!b3.ValidHash()){fmt.Println("ValidHash test6 passed")}
	*/
	/*
	//mining.go tests
	b0 := Initial(20)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	*/
	//blockchain.go tests
	//valid chain test
	/*
	b0 := Initial(20)
	b0.Mine(10)
	//fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(10)
	//fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(10)
	//fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	b3 := b2.Next("this is kinda interesting")
	b3.Mine(10)
	//fmt.Println(b3.Proof, hex.EncodeToString(b3.Hash))
	bChain := Blockchain{}
	bChain.Add(b0)
	bChain.Add(b1)
	bChain.Add(b2)
	bChain.Add(b3) 
	if(bChain.IsValid()){fmt.Println("test1: bChain is valid, pass")}
	//invalidate a block
	bChain.Chain[2].Difficulty = 1
	if(!bChain.IsValid()){fmt.Println("test2: bChain is invalid, pass")}
	bChain.Chain[2].Difficulty = 20 //reset chain
	if(bChain.IsValid()){fmt.Println("chain successfully reset")}
	badBlock := b2
	bChain.Chain = append(bChain.Chain, badBlock)
	if(!bChain.IsValid()){fmt.Println("test3: bChain is invalid, pass")}
	*/
	//Block Next() test
	b0 := Initial(16)
	//b1 := b0.Next("message")
	//Block CalcHash() test
	b0.SetProof(56231)
	//fmt.Printf("b0 hash is: %x\n", b0.Hash)
	bHash := hex.EncodeToString(b0.Hash)
	bString := "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000"
	fmt.Printf("bHash: %s, bString: %s\n", bHash, bString)
	bResult := string(bHash) != bString
	if(bResult){fmt.Println("they are equal")}
}