package main

import "fmt"

// Import the blockchain package

func main() {
	// Create a new blockchain instance using the NewBlockchain function.
	bc := NewBlockchain()

	// Add new blocks to the blockchain
	bc.AddBlock("YOOOO first block")
	bc.AddBlock("foooooobar 2nd block")

	// Iterate through each block in the blockchain and print block details.
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash) // Print the previous block's hash.
		fmt.Printf("Data: %s\n", block.Data)                // Print the data stored in the block.
		fmt.Printf("Hash: %x\n", block.Hash)                // Print the hash of the current block.
		fmt.Println()
	}
}
