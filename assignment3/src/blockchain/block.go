package blockchain

import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"
)

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