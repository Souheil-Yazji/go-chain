package blockchain

type Block struct {
	Index        int
	Timestamp    string
	Data         []byte // can be changed to a struct/other data types
	PreviousHash []byte
	Hash         []byte
}
