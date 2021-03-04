package main

import (
	"fmt"
	"github.com/boltdb/bolt-master"
	"os"
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

	db, err := bolt.Open(DBFILE,0600,nil)
	if err != nil {
		panic(err.Error())
	}

	blockchain := chain.CreatChainWithGenesis([]byte("hello"))

	blockchain.AddNewBlock([]byte("block1"))
	blockchain.AddNewBlock([]byte("block2"))
	fmt.Println("当前共有区块个数： ", len(blockchain.Blocks))
	//fmt.Println(blockchain.Blocks[0])
	//fmt.Println(blockchain.Blocks[1])
	//fmt.Println(blockchain.Blocks[2])

	block0 := blockchain.Blocks[0]
	block0SerBytes, err := block0.Serialize()
	if err != nil {
		fmt.Println("序列化区块0出现错误")
		return
	}
	deBlock0, err := chain.Deserialize(block0SerBytes)
	if err != nil {
		fmt.Println("反序列化区块0出现错误，程序已停止")
		return
	}
	fmt.Println(string(deBlock0.Data))


	//gensis := chain.CreateGenesisBlock([]byte("hello"))
	//fmt.Println("区块0: ", gensis)
	//block1 := chain.CreateBlock(gensis.Height,gensis.Hash,nil)
	//fmt.Println("区块1: ", block1)
	//block2 := chain.CreateBlock(gensis.Height,block1.Hash,nil)
	//fmt.Println("区块2: ", block2)

	bolt.Open()
	db := bolt.DB{}
	//读
	db.View(func(tx *bolt.Tx) error {
		return nil
	})
	//写
	db.Update(func(tx *bolt.Tx) error {
		return nil
	})
}


