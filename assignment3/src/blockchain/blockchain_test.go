package blockchain

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: some useful tests of Blocks
func TestQueueBasics(t *testing.T) {
	//block.go tests
	/*
	
	//Block Initial() test
	//fmt.Println(b0.PrevHash)
	assert.Equal(b0.PrevHash,true,"ERROR: invalid block PrevHash value")
	//fmt.Println(b0.Generation)
	assert.Equal(b0.Generation,true,"ERROR: invalid block Generation value")
	//fmt.Println(b0.Difficulty)
	assert.Equal(b0.Generation,true,"ERROR: invalid block Generation value")
	//fmt.Println(b0.PrevHash)
	assert.Equal(b0.Generation,true,"ERROR: invalid block Generation value")
	//fmt.Println(b0.Data)
	*/
	//Block Next() test
	b0 := Initial(16)
	b1 := b0.Next("message")
	//Block CalcHash() test
	b0.SetProof(56231)
	//fmt.Printf("b0 hash is: %x\n", b0.Hash)
	b0.Hash
	assert.Equal(hex.EncodeToString(b0.Hash),"6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000","ERROR: invalid block Hash value")
	b1 = b0.Next("message")
	b1.SetProof(2159)
	//fmt.Printf("b1 hash is: %x\n", b1.Hash)
	assert.Equal(hex.EncodeToString(b1.Hash),"9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000","ERROR: invalid block Hash value")
	/*
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
	//valid chain test
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
	//if(bChain.IsValid()){fmt.Println("test1: bChain is valid, pass")}
	assert.Equal(bChain.IsValid(),true,"ERROR: valid block chain failed validation")
	//invalidate a block
	bChain.Chain[2].Difficulty = 1
	if(!bChain.IsValid()){fmt.Println("test2: bChain is invalid, pass")}
	bChain.Chain[2].Difficulty = 20 //reset chain
	if(bChain.IsValid()){fmt.Println("chain successfully reset")}
	badBlock := b2
	bChain.Chain = append(bChain.Chain, badBlock)
	if(!bChain.IsValid()){fmt.Println("test3: bChain is invalid, pass")}
	*/	
}