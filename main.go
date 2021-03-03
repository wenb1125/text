package main

import (
	"fmt"
	"xianfengChain/chain"
)
/**
 * 步骤
 * 1.数据结构定义
 * 2.产生区块
 * 3.产生多个区块
 * 4.将区块连起来
		a.把已有的区块存起来
 * 5.把区块持久存起来
 * 6.多个节点直接同步数据
 */

/**
 * 项目的主入口
 */
func main() {
	fmt.Println("hello!!!")

	blockchain := chain.CreatChainWithGenesis([]byte("hello"))

	blockchain.AddNewBlock([]byte("block1"))
	blockchain.AddNewBlock([]byte("block2"))
	fmt.Println("当前共有区块个数： ", len(blockchain.Blocks))
	fmt.Println(blockchain.Blocks[0])
	fmt.Println(blockchain.Blocks[1])
	fmt.Println(blockchain.Blocks[2])



	//gensis := chain.CreateGenesisBlock([]byte("hello"))
	//fmt.Println("区块0: ", gensis)
	//block1 := chain.CreateBlock(gensis.Height,gensis.Hash,nil)
	//fmt.Println("区块1: ", block1)
	//block2 := chain.CreateBlock(gensis.Height,block1.Hash,nil)
	//fmt.Println("区块2: ", block2)
}


