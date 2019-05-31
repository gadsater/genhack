package chain

import (
	"main/lib"
	"encoding/json"
	"errors"
	"strconv"
)

type Chain struct {
	ChainLoc string
	ChainName string
	ChainId   string
	Index     string
	PrevHash  string
}

type ChainFunc interface {
	InitChain(string) (Chain, error)
	ReadChain(string) (Chain, error)
	IsChainExist(string) error
	JsonToChain([]byte) Chain
	ChainToJson(Chain) []byte
}
func (chain Chain) InitBlock() (Block, error) {
	if !lib.IsFileExist(chain.ChainLoc, chain.ChainName) {
		return Block{}, errors.New("Chain does not exist")
	}

	return Block{
		ChainName: chain.ChainName,
		Index:     chain.Index,
		Timestamp: lib.GetTime(),
		Sender:    "",
		Receiver:  "",
		Validator: "",
		Data:      "",
		PrevHash:  chain.PrevHash,
	}, nil
}

func (chain Chain) ReadBlock(blockHash string) (Block, error) {
	jsonBlock, err := lib.ReadFile(chain.ChainLoc, chain.ChainName, blockHash) 
	if err != nil {
		return Block{}, err
	}

	return JsonToBlock(jsonBlock), err
}

func (chain Chain) WriteBlock(block Block) error {
	jsonBlock := BlockToJson(block)

	blockHash := lib.GetShaString(string(jsonBlock))

	err := lib.WriteFile(chain.ChainLoc, chain.ChainName, blockHash, jsonBlock)
	if err != nil {
		return err
	}

	chain.PrevHash = blockHash
	chainInd, _ := strconv.Atoi(chain.Index)
	chain.Index = strconv.Itoa(chainInd + 1)

	err = lib.RmFile(chain.ChainLoc, chain.ChainName, chain.ChainId)
	if err != nil {
		return err
	}

	jsonChain := ChainToJson(chain)

	err = lib.WriteFile(chain.ChainLoc, chain.ChainName, chain.ChainId, jsonChain)
	
	return err
}


func ChainToJson(chain Chain) []byte {
	jsonChain, _ := json.MarshalIndent(chain, "", " ")
	return jsonChain
}

func JsonToChain(jsonData []byte) Chain {
	chain := Chain{}
	json.Unmarshal(jsonData, &chain)

	return chain
}
