package consensus

import (
	"math/big"
	"xianfengChain/chain"
	"xianfengChain/consensus/pos"
	"xianfengChain/consensus/pow"
)

/**
 * Consensus作为接口，用于定义共识机制的标准
 */
type Consensus interface {
	SearchNonce() ([32]byte,int64)
}

func NewProofWork(block chain.Block) Consensus {
	init := big.NewInt(1)
	init.Lsh(init,256 - pow.DIFFICULT)
	return pow.ProofWork{block,init}
}

func NewStock() Consensus {
	return pos.ProofStock{}
}