package pos

import (
	"fmt"
	"xianfengChain/chain"
)

/**
 *
 */
type ProofStock struct {
	Block chain.Block
}


func (stock ProofStock) SearchNonce() ([32]byte,int64){
	fmt.Println("我是新来的，这个是我写的共识机制的pos的实现方法")
	return [32]byte{},0
}