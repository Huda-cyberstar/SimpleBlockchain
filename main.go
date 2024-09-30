package main

import (
	"assignment01bca"
	"fmt"
)

func main() {
	// Creating the Genesis block
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	fmt.Println("Genesis block created")

	// Adding more blocks
	assignment01bca.NewBlock("Alice pays Bob 10 BTC", 1, genesisBlock.Hash)
	assignment01bca.NewBlock("Bob pays Charlie 5 BTC", 2, assignment01bca.Blockchain[1].Hash)

	// List all blocks
	fmt.Println("\nListing all blocks in the blockchain:")
	assignment01bca.ListBlocks()

	// Change a block transaction
	fmt.Println("\nModifying the second block's transaction:")
	assignment01bca.ChangeBlock(1, "Alice pays Bob 15 BTC")

	// Verify the blockchain after modification
	fmt.Println("\nVerifying the blockchain:")
	assignment01bca.VerifyChain()

	// List all blocks after modification
	fmt.Println("\nListing all blocks in the blockchain after modification:")
	assignment01bca.ListBlocks()
}
