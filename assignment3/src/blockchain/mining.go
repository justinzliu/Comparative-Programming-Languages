package blockchain

import (
	"work_queue"
)

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

