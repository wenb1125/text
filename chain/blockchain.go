package chain

import "github.com/boltdb/bolt-master"

const BLOCKS = "blocks"
const LASTHASH = "lasthash"

/**
 * 定义区块链结构体，用于存储产生的区块（内存中）
 */
type BlockChain struct {
	//Blocks []Block
	//文件操作对象
	DB     *bolt.DB
}

//func NewBlockChain() BlockChain {
//	return BlockChain{}
//}

func NewBlockChain(db *bolt.DB) BlockChain {
	return BlockChain{db}
}

/**
 * 创建一个区块链实例，该实例携带一个创世区块
 */
func (chain *BlockChain) CreatChainWithGenesis(genesisData []byte) BlockChain {
	genesis := CreateGenesisBlock(genesisData)
	genSerBytes, _ := genesis.Serialize()

	db := chain.DB

	//读一遍bucket，查看是否有数据
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil {//第一次为nil
			//程序第一次执行的写入数据
			db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte(BLOCKS))
				if bucket == nil {
					bucket, _ = tx.CreateBucket([]byte(BLOCKS))
				}
				bucket.Put(genesis.Hash[:],genSerBytes)
				return nil
			})
			return nil
		}
		blockBytes := bucket.Get([]byte(LASTHASH))
		if len(blockBytes) == 0 {
			//程序第一次执行的写入数据
			db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte(BLOCKS))
				if bucket == nil {
					bucket, _ = tx.CreateBucket([]byte(BLOCKS))
				}
				bucket.Put(genesis.Hash[:],genSerBytes)
				return nil
			})
			return nil
		}

		return nil
	})



	blocks := make([]Block, 0)
	blocks = append(blocks, genesis)
	return BlockChain{blocks}
}

func (chain *BlockChain) AddNewBlock(data []byte) {
	//1.找到切片的最后一个元素，代表最新的区块
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	//2.根据最后一个区块产生一个新区块
	newBlock := CreateBlock(lastBlock.Height, lastBlock.Hash, data)
	//3.把最新产生的区块放入到切片中
	chain.Blocks = append(chain.Blocks, newBlock)
}
