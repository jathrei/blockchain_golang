package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Timestamp     int64  //current time stamp at time of block creation
	Data          []byte //actual data stored in block
	PrevBlockHash []byte //stores the hash of previous block
	Hash          []byte //hash of current block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) //taking the timestamp int64 and converting to string, then converting that string
	//to a slice of bytes
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	//this code above creates a slice of slice of type []byte of b.PrevBlockHash,b.Data,timestamp.. and then concatenates them with an empty
	//seperator into a single slice of bytes "[]byte" into header (header of type []byte)
	hash := sha256.Sum256(headers)
	//takes the hash value of the headers slice of bytes and creates a [32]byte hash value (specifically, 256 bits or 32 bytes)
	b.Hash = hash[:]
	//set the b.Hash value to the hash value, but convert it from an array to a slice (by doing [:]).. purpose of this is for consistency
	//plus future utilization of the slice

	//remember, since we passed by reference, changing b.Hash changes the value outside of the scope of the func
}

// function to create a new block is passed in data
func NewBlock(data string, prevBlockHash []byte) *Block {

}
