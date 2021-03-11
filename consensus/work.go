package consensus

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"xianfengChain/utils"
)

//256位二进制
//思路：给一个大整数，初始值为1，根据自己需要的难度进行左移位，左移位的位数是256-0的个数

const DIFFICULTY = 16 //初始难度为10，即大整数到底开头有10个0

/**
 * 工作量证明
 */
type ProofWork struct {
	Block  BlockInterface
	Target *big.Int
}

/**
 * 实现共识机制接口的方法
 */
func (work ProofWork) SearchNonce() ([32]byte,int64) {
	//block -> nonce
	//block哈希 < 系统提供的某个目标值
	//1.给定一个nonce值,重新计算带有nonce的区块哈希
	var nonce int64
	nonce = 0
	hashBig := new(big.Int)
	for {
		hash := CalculateBlockHash(work.Block, nonce)
		//2.系统给定的值
		target := work.Target
		//3.拿1和2比较
		hashBig := hashBig.SetBytes(hash[:])
		result := hashBig.Cmp(target)
		//result := bytes.Compare(hash[:], target.Bytes())
		//4.判断结果,区块哈希 < 给定值, 返回nonce; 否则nonce自增
		if result == -1 {
			return hash, nonce
		}
		nonce++
	}
}

func CalculateBlockHash(block BlockInterface, nonce int64) [32]byte {
	heightByte, _ := utils.Int2Byte(block.GetHeight())
	versionByte, _ := utils.Int2Byte(block.GetVersion())
	timeByte, _ := utils.Int2Byte(block.GetTimeStamp())
	nonceByte, _ := utils.Int2Byte(nonce)
	preHash := block.GetPreHash()
	bk := bytes.Join([][]byte{
		heightByte,
		versionByte,
		preHash[:],
		timeByte,
		nonceByte,
		},
		[]byte{})
	return sha256.Sum256(bk)
}
