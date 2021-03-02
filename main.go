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
	fmt.Println("新区块: ", gensis)
}
