package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Block struct {
	BlockchainId string
	Index        int
	Timestamp    int64
	Sender       string
	Receiver     string
	Validator    string
	Data         string
	PrevHash     string
}

type BlockchainInfo struct {
	BlockchainId string
	Index        int
	PrevHash     string
}

func InitBlock(_bcid string, _index int, _data string, _prev_hash string) Block {
	return Block{
		BlockchainId: _bcid,
		Index:        _index,
		Timestamp:    time.Now().Unix(),
		Sender:       "",
		Receiver:     "",
		Validator:    "",
		Data:         _data,
		PrevHash:     _prev_hash,
	}
}

func WriteBlock(BlockchainDir string, _block Block) BlockchainInfo {
	json_block, err := json.MarshalIndent(_block, "", " ")
	check(err)

	sha_file := sha256.Sum256(json_block)

	FileName := hex.EncodeToString(sha_file[:])
	err = ioutil.WriteFile(
		BlockchainDir+FileName,
		json_block,
		0444)
	check(err)

	return BlockchainInfo{
		BlockchainId: _block.BlockchainId,
		Index:        _block.Index + 1,
		PrevHash:     FileName,
	}
}

func readBlockchain(file_loc string, blockchain_id string) BlockchainInfo {
	data, err := ioutil.ReadFile(file_loc + "/" + blockchain_id)

	block_info := BlockchainInfo{}
	err = json.Unmarshal(data, &block_info)

	if err != nil {
		block_info = BlockchainInfo{
			BlockchainId: "",
			Index:        0,
			PrevHash:     "",
		}
	}
	return block_info
}

func writeBlockchain(file_loc string, blockchain_id string, data BlockchainInfo) {
	err := os.Remove(file_loc + "/" + blockchain_id)

	block_info, err := json.MarshalIndent(data, "", " ")

	err = ioutil.WriteFile(file_loc+"/"+blockchain_id, block_info, 0444)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	blockchain_id := "d439c8e6552b43c2bd3ccc759db542ac5778378e662a5bcdd84e7cb222ef51c7"

	prev_block := readBlockchain("./BlockDir", blockchain_id)

	prev_block = WriteBlock("./BlockDir/",
		InitBlock(blockchain_id, prev_block.Index,
			"Hello", prev_block.PrevHash))

	writeBlockchain("./BlockDir", blockchain_id, prev_block)
}
