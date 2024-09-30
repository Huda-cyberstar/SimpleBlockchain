package main

import (
	"blockchainproject/assignment01bca" // Import the package
	"fmt"
)

func main() {
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	fmt.Println("Genesis block created")

	assignment01bca.NewBlock("Alice pays Bob 10 BTC", 1, genesisBlock.Hash)
	assignment01bca.NewBlock("Bob pays Charlie 5 BTC", 2, assignment01bca.Blockchain[1].Hash)

	fmt.Println("\nListing all blocks in the blockchain:")
	assignment01bca.ListBlocks()

	fmt.Println("\nModifying the second block's transaction:")
	assignment01bca.ChangeBlock(1, "Alice pays Bob 15 BTC")

	fmt.Println("\nVerifying the blockchain:")
	if assignment01bca.VerifyChain() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}

	fmt.Println("\nListing all blocks in the blockchain after modification:")
	assignment01bca.ListBlocks()
}
