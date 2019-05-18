package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Block struct {
	Index     int
	Timestamp int64
	Sender    []byte
	Receiver  []byte
	Validator []byte
	Data      string
	Prev_hash []byte
}

type BlockchainInfo struct {
	Index     int
	Prev_hash []byte
}

func InitBlock(_index int, _data string, _prev_hash []byte) Block {
	return Block{
		Index:     _index,
		Timestamp: time.Now().Unix(),
		Sender:    nil,
		Receiver:  nil,
		Validator: nil,
		Data:      _data,
		Prev_hash: _prev_hash,
	}
}

func WriteBlock(BlockchainDir string, _block Block) BlockchainInfo {
	json_block, err := json.MarshalIndent(_block, "", " ")
	check(err)

	shaFile := sha1.New()
	shaFile.Write(json_block)

	FileName := hex.EncodeToString(shaFile.Sum(nil))
	err = ioutil.WriteFile(
		BlockchainDir+FileName,
		json_block,
		0444)
	check(err)

	return BlockchainInfo{
		Index:     _block.Index + 1,
		Prev_hash: []byte(FileName),
	}
}

func readPrevDetail(file_loc string) BlockchainInfo {

	data, err := ioutil.ReadFile(file_loc)

	block_info := BlockchainInfo{}
	err = json.Unmarshal(data, &block_info)

	if err != nil {
		block_info = BlockchainInfo{
			Index:     0,
			Prev_hash: nil,
		}
	}
	return block_info
}

func writePrevDetail(file_loc string, data BlockchainInfo) {
	err := os.Remove(file_loc)

	block_info, err := json.MarshalIndent(data, "", " ")

	err = ioutil.WriteFile(file_loc, block_info, 0444)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	prev_block := readPrevDetail("./BlockDir/prev_block")

	prev_block = WriteBlock("./BlockDir/",
		InitBlock(prev_block.Index, "Hello", prev_block.Prev_hash))

	writePrevDetail("./BlockDir/prev_block", prev_block)
}
