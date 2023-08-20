package main

// It holds a collection of blocks that form the blockchain.
type Blockchain struct {
	blocks []*Block // A slice to store pointers to individual blocks in the blockchain.
}

// adds a new block to the blockchain w/ data
func (bc *Blockchain) AddBlock(data string) {
	// Get a reference to the previous block in the blockchain.
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new block using the NewBlock function
	//pass in the data and previous block's hash.
	newBlock := NewBlock(data, prevBlock.Hash)

	// Append the newly created block to the blocks slice in the blockchain.
	bc.blocks = append(bc.blocks, newBlock)
}

func NewFirstBlock() *Block {
	// Create a new block using the NewBlock function, passing in the data and an empty previous hash.
	return NewBlock("First Block", []byte{})
}

// NewBlockchain creates and returns a new instance of the Blockchain struct, initialized with the first block.
func NewBlockchain() *Blockchain {
	// Create the first block of the blockchain using NewFirstBlock function.
	firstBlock := NewFirstBlock()

	// Create a new Blockchain instance with a slice containing the first block.
	blockchain := &Blockchain{
		blocks: []*Block{firstBlock},
	}

	return blockchain
}
