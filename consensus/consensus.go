package consensus

import (
	"math/big"
)

/**
 * Consensus作为接口，用于定义共识机制的标准
 */
type Consensus interface {
	SearchNonce() ([32]byte, int64)
}

type BlockInterface interface {
	GetHeight() int64
	GetVersion() int64
	GetTimeStamp() int64
	GetPreHash() [32]byte
	GetData() []byte
}

func NewProofWork(block BlockInterface) Consensus {
	init := big.NewInt(1)
	init.Lsh(init, 256 - DIFFICULTY)
	return ProofWork{block, init}
}

func NewStock() Consensus {
	return ProofStock{}
}
