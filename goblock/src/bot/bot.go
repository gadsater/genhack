package bot

import(
	"main/lib"
	"chain"
)

type Bot struct {
	BotName string
	FileLoc string
}

func InitBot(botName string, fileLoc string) Bot {
	return Bot {
		BotName: botName,
		FileLoc: fileLoc,
	}
}

func (b Bot) InitChain(chainName string) (chain.Chain, error) {

	if b.IsChainExist(chainName) {
		chain, err := b.ReadChain(chainName)
		return chain, err
	} else {
		err := lib.MakeDir(b.FileLoc, chainName)
		if err != nil {
			return chain.Chain{}, err
		}
		
		chainId := lib.GetShaString(chainName)

		chainData := chain.Chain{
			ChainLoc: b.FileLoc,
			ChainName: chainName,
			ChainId:   chainId,
			Index:     "0",
			PrevHash:  "",
		}

		jsonChain := chain.ChainToJson(chainData)

		err = lib.WriteFile(b.FileLoc, chainName, chainId, jsonChain)
		if err != nil {
			return chain.Chain{}, err
		}

		return chainData, err
	}
}

func (b Bot) ReadChain(chainName string) (chain.Chain, error) {
	chainId := lib.GetShaString(chainName)

	jsonChain, err := lib.ReadFile(b.FileLoc, chainName, chainId)
	if err != nil {
		return chain.Chain{}, err
	}

	chainData := chain.JsonToChain(jsonChain)

	return chainData, nil
}

func (b Bot) IsChainExist(chainName string) bool {
	return lib.IsFileExist(b.FileLoc, chainName)
}

func check(e error) {
	if e != nil {
		panic(e);
	}
}
