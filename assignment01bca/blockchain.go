package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

// CalculateHash calculates the hash for a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	return hex.EncodeToString(hash.Sum(nil))
}

// NewBlock creates a new block and adds it to the blockchain
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(transaction + string(nonce) + previousHash)
	Blockchain = append(Blockchain, *block)
	return block
}

// ListBlocks lists all the blocks in the blockchain
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("-------------------------------")
	}
}

// ChangeBlock changes the transaction of a block at the given index
func ChangeBlock(index int, newTransaction string) {
	if index < len(Blockchain) {
		Blockchain[index].Transaction = newTransaction
		Blockchain[index].Hash = CalculateHash(newTransaction + string(Blockchain[index].Nonce) + Blockchain[index].PreviousHash)
		fmt.Printf("Block %d transaction changed to: %s\n", index+1, newTransaction)
	} else {
		fmt.Println("Invalid block index")
	}
}

// VerifyChain checks the integrity of the blockchain
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		previousBlock := Blockchain[i-1]
		currentBlock := Blockchain[i]
		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Blockchain compromised at block %d!\n", i+1)
			return false
		}
	}
	fmt.Println("Blockchain is valid.")
	return true
}
