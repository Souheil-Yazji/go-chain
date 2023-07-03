package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"go-chain/pkg/blockchain"
	"time"
)

var Blockchain []blockchain.Block

func main() {

}

// init a block
func createBlk(data string) blockchain.Block {
	return blockchain.Block{
		Index:        0,
		Timestamp:    "",
		Data:         []byte(data),
		PreviousHash: nil,
		Hash:         nil,
	}
}

// dataFmt: 0 for byte slice, 1 for string representation
func printBlk(block blockchain.Block, dataFmt int) {
	fmt.Print(String(block, dataFmt))
}

func calcHash(block blockchain.Block) []byte {
	record := fmt.Sprintf("%d%s%s", block.Index, block.Timestamp, block.Data)
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hash
}

// the byte array output can be useful for debugging
func String(block blockchain.Block, dataFmt int) string {
	str := ""
	if dataFmt == 0 {
		str = fmt.Sprintf("Block: %+v\n", block) // %+v gives field name + default value
	} else if dataFmt == 1 {
		str = fmt.Sprintf("Block: {Index:%d Timestamp:%s Data:%q PreviousHash:%v Hash:%v}\n",
			block.Index, block.Timestamp, block.Data, block.PreviousHash, block.Hash)
	}
	return str
}

// generate the next block in the chain
func genBlk(prevBlock blockchain.Block, data string) (blockchain.Block, error) {
	var newBlock blockchain.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = []byte(data)
	newBlock.PreviousHash = prevBlock.Hash
	newBlock.Hash = calcHash(newBlock)

	return newBlock, nil
}

// validate block prior to commiting it to the chain
func validateBlk(block blockchain.Block) bool {
	calculatedHash := calcHash(block)

	if !bytes.Equal(block.Hash, calculatedHash) {
		return false
	}

	if block.Index > 0 {
		previousBlock := Blockchain[block.Index-1]
		if !bytes.Equal(previousBlock.Hash, block.PreviousHash) {
			return false
		}
	}

	return true
}
