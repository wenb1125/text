package main

import (
	"fmt"
	"github.com/boltdb/bolt-master"
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

const DBFILE = "xianfneg.db"

/**
 * 项目的主入口
 */
func main() {
	fmt.Println("hello!!!")

	engine, err := bolt.Open(DBFILE,0600,nil)
	if err != nil {
		panic(err.Error())
	}

	blockChain := chain.NewBlockChain(engine)
	//创世区块
	blockChain.CreatGenesis([]byte("hello word"))
	//新增一个区块
	err = blockChain.AddNewBlock([]byte("hello"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//获取区块
	//lastBlock, err := blockChain.GetLastBlock()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(lastBlock)

	allBlocks, err := blockChain.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, block := range allBlocks {
		fmt.Println(block)
	}
}


