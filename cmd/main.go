package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	Blockchain "go-chain/pkg/blockchain"
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

var blockchain []Blockchain.Block

func main() {

}

// init a block
func createBlk(data string) Blockchain.Block {
	return Blockchain.Block{
		Index:        0,
		Timestamp:    "",
		Data:         []byte(data),
		PreviousHash: nil,
		Hash:         nil,
	}
}

// dataFmt: 0 for byte slice, 1 for string representation
func printBlk(block Blockchain.Block, dataFmt int) {
	fmt.Print(String(block, dataFmt))
}

func calcHash(block Blockchain.Block) []byte {
	record := fmt.Sprintf("%d%s%s", block.Index, block.Timestamp, block.Data)
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hash
}

// the byte array output can be useful for debugging
func String(block Blockchain.Block, dataFmt int) string {
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
func genBlk(prevBlock Blockchain.Block, data string) (blockchain.Block, error) {
	var newBlock Blockchain.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = []byte(data)
	newBlock.PreviousHash = prevBlock.Hash
	newBlock.Hash = calcHash(newBlock)

	return newBlock, nil
}

// validate block prior to commiting it to the chain
func validateBlk(block Blockchain.Block) bool {
	calculatedHash := calcHash(block)

	if !bytes.Equal(block.Hash, calculatedHash) {
		return false
	}

	if block.Index > 0 {
		previousBlock := blockchain[block.Index-1]
		if !bytes.Equal(previousBlock.Hash, block.PreviousHash) {
			return false
		}
	}

	return true
}

func openDatabase() (*badger.DB, error) {
	options := badger.DefaultOptions("").WithInMemory(false)
	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}
	return db, nil
}
