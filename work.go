package main

import "math/big"

//proof of work

// targetBits is a constant that represents the difficulty level for mining
// It determines the number of leading zeros required in the hash of a block
const targetBits = 16

// ProofOfWork represents the proof of work algorithm for mining a block
type ProofOfWork struct {
	block  *Block   // The block for which proof of work is being calculated
	target *big.Int // The target value miners need to find a hash below
}

// NewProofOfWork creates and returns a new ProofOfWork
// It takes a block as input and calculates the target value based on the targetBits
func NewProofOfWork(b *Block) *ProofOfWork {
	// Create a new big.Int instance initialized with 1
	target := big.NewInt(1)

	// Left shift the target value by (256 - targetBits) positions
	// This creates a target value with the desired number of trailing zeros
	target.Lsh(target, uint(256-targetBits))

	// Create a new ProofOfWork instance with the calculated target value and the given block
	work := &ProofOfWork{b, target}

	return work
}
