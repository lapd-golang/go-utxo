package blockchain

import "math/big"

const targetBits = 24

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewPOW(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}
