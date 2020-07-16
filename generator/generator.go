package generator

import (
	"crypto/rand"
	"math/big"
)

type Generator interface {
	Generate() ([]byte, error)
}

type BytesGenerator int

func (g BytesGenerator) Generate() ([]byte, error) {
	buf := make([]byte, int(g))
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

type ListGenerator struct {
	List []byte
	Min  int
	Max  int
}

func (g *ListGenerator) Generate() ([]byte, error) {
	var length int
	if g.Max <= g.Min {
		length = int(g.Min)
	} else {
		max, err := rand.Int(rand.Reader, big.NewInt(int64(g.Max-g.Min)))
		if err != nil {
			return nil, err
		}
		length = g.Min + int(max.Int64())
	}
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(g.List))))
		if err != nil {
			return nil, err
		}
		result[i] = g.List[int(index.Int64())]
	}
	return result, nil
}
