package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Block struct {
	ChainName string
	Index     int
	Timestamp int64
	Sender    string
	Receiver  string
	Validator string
	Data      string
	PrevHash  string
}

type ChainInfo struct {
	ChainName string
	ChainId   string
	Index     int
	PrevHash  string
}

func initBlock(fileLoc string, chainName string) (Block, error) {
	if _, err := os.Stat(fileLoc + "/" + chainName); os.IsNotExist(err) {
		fmt.Println("Chain doesn't exist, please init chain first. Hint: Refer Initchain function")
		return Block{}, err
	}

	chainInfo, _ := ReadChain(fileLoc, chainName)

	return Block{
		ChainName: chainName,
		Index:     chainInfo.Index,
		Timestamp: time.Now().Unix(),
		Sender:    "",
		Receiver:  "",
		Validator: "",
		Data:      "",
		PrevHash:  chainInfo.PrevHash,
	}, nil
}

func ReadBlock(fileLoc string, chainId string, blockHash string) (Block, error) {
	jsonBlock, err := ioutil.ReadFile(fileLoc + "/" + chainId + "/" + blockHash)
	if err != nil {
		fmt.Println("Requested block doesn't exist")
		return Block{}, err
	}

	block := Block{}
	err = json.Unmarshal(jsonBlock, &block)

	return block, nil
}

func WriteBlock(fileLoc string, chainName string, data string) error {
	block, err := initBlock(fileLoc, chainName)
	block.Data = data

	jsonBlock, err := json.MarshalIndent(block, "", " ")
	check(err)

	blockHash := getShaString(string(jsonBlock))

	err = ioutil.WriteFile(
		fileLoc+"/"+block.ChainName+"/"+blockHash,
		jsonBlock,
		0444)

	if err != nil {
		fmt.Println("Blockchain is unintialized")
		return err
	}

	err = updateChain(fileLoc, block.ChainName, blockHash)

	return nil
}

func getShaString(chainName string) string {
	chainHash := sha256.Sum256([]byte(chainName))
	return hex.EncodeToString(chainHash[:])
}

func InitChain(fileLoc string, chainName string) (ChainInfo, error) {

	if _, err := os.Stat(fileLoc + "/" + chainName); os.IsNotExist(err) {

		err = os.Mkdir(fileLoc+"/"+chainName, 0777)

		if err != nil {
			fmt.Println("Problem initializing the chain.")
			return ChainInfo{}, err
		}

		chainHash := sha256.Sum256([]byte(chainName))
		chainId := hex.EncodeToString(chainHash[:])

		chainInfo := ChainInfo{
			ChainName: chainName,
			ChainId:   chainId,
			Index:     0,
			PrevHash:  "",
		}

		jsonChain, _ := json.MarshalIndent(chainInfo, "", " ")

		err = ioutil.WriteFile(fileLoc+"/"+chainName+"/"+chainId, jsonChain, 0444)

		return chainInfo, nil

	} else {

		chainInfo, _ := ReadChain(fileLoc, chainName)
		return chainInfo, nil

	}

}

func ReadChain(fileLoc string, chainName string) (ChainInfo, error) {
	chainId := getShaString(chainName)

	jsonChain, err := ioutil.ReadFile(fileLoc + "/" + chainName + "/" + chainId)

	if err != nil {
		fmt.Println("Chain doesn't exist")
		return ChainInfo{}, err
	}

	chainInfo := ChainInfo{}
	json.Unmarshal(jsonChain, &chainInfo)

	return chainInfo, nil
}

func updateChain(fileLoc string, chainName string, prevHash string) error {
	chainId := getShaString(chainName)

	chainInfo, err := ReadChain(fileLoc, chainName)

	if err != nil {
		fmt.Println("Error updating chain; chain may not exist")
		return err
	}

	chainInfo.Index = chainInfo.Index + 1
	chainInfo.PrevHash = prevHash

	err = os.Remove(fileLoc + "/" + chainName + "/" + chainId)

	jsonChain, err := json.MarshalIndent(chainInfo, "", " ")

	err = ioutil.WriteFile(fileLoc+"/"+chainName+"/"+chainId, jsonChain, 0444)
	check(err)

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	chainName := "TornadoTurtle"
	chainLoc := "../store/chains"

	_, err := InitChain(chainLoc, chainName)
	check(err)

	_ = WriteBlock(chainLoc, chainName, "Hello World")

}
