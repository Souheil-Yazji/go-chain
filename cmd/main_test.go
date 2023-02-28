package main

// *testing.T allows output writes to appropriate test stream
import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-chain/pkg/blockchain"
)

var testData = "Hello, World!"

func TestCreateBlk(t *testing.T) {
	block := createBlk(testData)
	assert.Equal(t, 0, block.Index)
	assert.Equal(t, 0, block.Timestamp)
	assert.Equal(t, []byte(testData), block.Data)
	assert.Nil(t, block.PreviousHash)
	assert.Nil(t, block.Hash)
}

func TestCalcHash(t *testing.T) {
	block := blockchain.Block{
		Index:        0,
		Timestamp:    0,
		Data:         []byte(testData),
		PreviousHash: nil, // same as []byte{}
		Hash:         nil,
	}
	hash := calcHash(block)
	expectedHash := sha256.Sum256([]byte("0" + "0" + testData))
	assert.Equal(t, expectedHash[:], hash)
}

func TestString(t *testing.T) {
	byteData := 0
	strData := 1
	block := blockchain.Block{
		Index:        0,
		Timestamp:    0,
		Data:         []byte(testData),
		PreviousHash: nil,
		Hash:         nil,
	}

	// Test byte slice data represenation
	assert.Equal(
		t,
		"Block: {Index:0 "+
			"Timestamp:0 "+
			"Data:[72 101 108 108 111 44 32 87 111 114 108 100 33] "+
			"PreviousHash:[] "+
			"Hash:[]}\n",
		String(block, byteData))

	// Test str data represetation
	assert.Equal(
		t,
		"Block: {Index:0 "+
			"Timestamp:0 "+
			"Data:\"Hello, World!\" "+
			"PreviousHash:[] "+
			"Hash:[]}\n",
		String(block, strData))
}
