package main

import (
	"bot"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	chainName := "Sid"
	chainLoc := "../store/chains"

	druid := bot.InitBot("Druid", chainLoc)

	torTut, err := druid.InitChain(chainName)
	
	bl1, err := torTut.InitBlock()
	check(err)

	bl1.BlockAddData("hello")
	bl1.BlockAddData("world")
	bl1.BlockAddData("data")
	torTut.WriteBlock(bl1)
}
