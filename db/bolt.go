package db

import (
	"fmt"
	"github.com/boltdb/bolt-master"
	"xianfengChain/chain"
)

//自定义的存储引擎，用于实现对区块数据的读写操作
type DBEngine struct {
	DB *bolt.DB
}

//区块存入文件（写）
func (engine DBEngine) SaveBlock2DB(block chain.Block) {
	fmt.Println("在该方法中将区块存到db中去")
}

//从文件中恢复（读）
func (engine DBEngine) GetBlockFromDB(hash [32]byte) chain.Block {
	fmt.Println("该方法中从db中获取特定的区块")
	return chain.Block{}
}
