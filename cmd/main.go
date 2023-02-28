package main

import (
	"crypto/sha256"
	"fmt"
	"go-chain/pkg/blockchain"
)

var Blockchain []blockchain.Block

func main() {

}

func createBlk(data string) blockchain.Block {
	return blockchain.Block{
		Index:        0,
		Timestamp:    0,
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
	record := fmt.Sprintf("%d%d%s", block.Index, block.Timestamp, block.Data)
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hash
}

func String(block blockchain.Block, dataFmt int) string {
	str := ""
	if dataFmt == 0 {
		str = fmt.Sprintf("Block: %+v\n", block) // %+v gives field name + value
	} else if dataFmt == 1 {
		str = fmt.Sprintf("Block: {Index:%d Timestamp:%d Data:%q PreviousHash:%v Hash:%v}\n",
			block.Index, block.Timestamp, block.Data, block.PreviousHash, block.Hash)
	}
	return str
}
