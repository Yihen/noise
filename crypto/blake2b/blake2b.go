package blake2b

import (
	"github.com/perlin-network/noise/crypto"
	"golang.org/x/crypto/blake2b"
)

// Blake2b represents the BLAKE2 cryptographic hash algorithm.
type Blake2b struct{}

var (
	_ crypto.HashPolicy = (*Blake2b)(nil)
)

// New returns a BLAKE2 hash policy.
func New() *Blake2b {
	return &Blake2b{}
}

// HashBytes hashes the given bytes using the BLAKE2 hash algorithm.
func (p *Blake2b) HashBytes(bytes []byte) []byte {
	result := blake2b.Sum256(bytes)
	return result[:]
}
