package blockchain

import(
	"bytes"
)

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
