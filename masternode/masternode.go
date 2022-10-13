package masternode

import (
	"golang.org/x/crypto/sha3"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// MasterNodePing is an masternode ping.
type MasterNodePing struct {
	version int             `json:"version"       gencodec:"required"`
	signTime time.Time      `json:"signTime"       gencodec:"required"`
	sign []byte             `json:"sign"       gencodec:"required"`
	blockhash common.Hash   `json:"blockhash"       gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}
// NewMasetrNodePing creates a new masternode ping.
func NewMasetrNodePing() *MasterNodePing {
	mnp := new(MasterNodePing)
	return mnp
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Hash returns the transaction hash.
func (mnp *MasterNodePing) Hash() common.Hash {
	if hash := mnp.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	h := rlpHash(mnp)
	mnp.hash.Store(h)
	return h
}
