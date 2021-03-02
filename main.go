package main

import (
	"fmt"
	"xianfengChain/chain"
)

/**
 * 项目的主入口
 */
func main() {
	fmt.Println("hello!!!")

	gensis := chain.CreateGenesisBlock([]byte("hello"))
	fmt.Println("区块0: ", gensis)
	block1 := chain.CreateBlock(gensis.Height,gensis.Hash,nil)
	fmt.Println("区块1: ", block1)
	block2 := chain.CreateBlock(gensis.Height,block1.Hash,nil)
	fmt.Println("区块2: ", block2)
}


