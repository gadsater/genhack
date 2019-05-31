package chain

import (
	"main/lib"
	"encoding/json"
)

type Block struct {
	ChainName string `json:"chainName"`
	Index     string `json:"index"`
	Timestamp string `json:"timeStamp"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Validator string `json:"validator"`
	Data      string `json:"data"`
	PrevHash  string `json:"prevHash"`
}

type BlockFunc interface {
	InitBlock(string) (Block, error)
	ReadBlock(string, string) (Block, error)
	BlockAddData(string) error
	WriteBlock(Block) error
	JsonToBlock([]byte) Block
	BlockToJson(Chain) []byte
}

func (bl *Block) BlockAddData(data string) error {
	bl.Data = bl.Data + "{" + lib.GetTime() + " | " + data + "}" + "\n"
	return nil
}

func BlockToJson(block Block) []byte {
	jsonBlock, _ := json.MarshalIndent(block, "", " ")
	return jsonBlock
}

func JsonToBlock(jsonData []byte) Block {
	block := Block{}
	json.Unmarshal(jsonData, &block)

	return block
}
