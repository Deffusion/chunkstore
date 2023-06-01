package chunker

import (
	"github.com/whyrusleeping/chunker"
	"hash"
	"io"
)

type Rabin struct {
	reader  io.Reader
	hash    hash.Hash
	chunker *chunker.Chunker
}

func NewRabin(r io.Reader, h hash.Hash, avgBlkSize uint64) *Rabin {
	min := avgBlkSize / 3
	max := avgBlkSize + (avgBlkSize / 2)

	poly, _ := chunker.RandomPolynomial()
	c := chunker.New(r, poly, h, avgBlkSize, min, max)
	return &Rabin{
		r,
		h,
		c,
	}
}

func (r *Rabin) NextChunk() (*BasicChunk, error) {
	chunk, err := r.chunker.Next()
	if err != nil {
		return nil, err
	}
	return &BasicChunk{chunk}, nil
}