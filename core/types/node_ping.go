package types

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math/big"
	"sync/atomic"
)

const NodePingVersion int = 1002

const (
	MasterNodeType int = iota + 1
	SuperNodeType
)

type NodePingBase struct {
	Version   *big.Int
	Id        *big.Int
	NodeType  *big.Int
	CurBlock  common.Hash
	CurHeight *big.Int
	Time      *big.Int
}

type NodePing struct {
	Version   *big.Int     `json:"version"`
	Id        *big.Int     `json:"id"`
	NodeType  *big.Int     `json:"nodeType"`
	CurBlock  common.Hash  `json:"curBlock"`
	CurHeight *big.Int     `json:"curHeight"`
	Time      *big.Int     `json:"time"`
	V         *big.Int     `json:"v"`
	R         *big.Int     `json:"r"`
	S         *big.Int     `json:"s"`

	// caches
	hash atomic.Value
	size atomic.Value
}

func NewNodePing(id *big.Int, nodeType int, blockHash common.Hash, height *big.Int, privateKey *ecdsa.PrivateKey) (*NodePing, error) {
	base := &NodePingBase{ big.NewInt(int64(NodePingVersion)), id, big.NewInt(int64(nodeType)), blockHash, height, common.Big0}
	h := rlpHash(base)
	sig, err := crypto.Sign(h.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	r, s, v := decodeSignature(sig)
	return &NodePing{Version: base.Version, Id: base.Id, NodeType: base.NodeType, CurBlock: base.CurBlock, CurHeight: base.CurHeight, Time: base.Time, V: v, R: r, S: s}, nil
}

// Hash returns the transaction hash.
func (ping *NodePing) Hash() common.Hash {
	if hash := ping.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	h := rlpHash([]interface{}{
		ping.Version,
		ping.Id,
		ping.NodeType,
		ping.CurBlock,
		ping.CurHeight,
		ping.Time,
	})
	ping.hash.Store(h)
	return h
}

func (ping *NodePing) Size() common.StorageSize {
	if size := ping.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, ping)
	ping.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

type extNodePing struct {
	Version   *big.Int
	Id        *big.Int
	NodeType  *big.Int
	CurBlock  common.Hash
	CurHeight *big.Int
	Time      *big.Int
	V, R, S   *big.Int
}

// DecodeRLP decodes the Ethereum
func (ping *NodePing) DecodeRLP(s *rlp.Stream) error {
	var extPing extNodePing
	_, size, _ := s.Kind()
	if err := s.Decode(&extPing); err != nil {
		return err
	}
	ping.Version = extPing.Version
	ping.Id = extPing.Id
	ping.NodeType = extPing.NodeType
	ping.CurBlock = extPing.CurBlock
	ping.CurHeight = extPing.CurHeight
	ping.Time = extPing.Time
	ping.V = extPing.V
	ping.R = extPing.R
	ping.S = extPing.S
	ping.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the Ethereum RLP block format.
func (ping *NodePing) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extNodePing{ping.Version, ping.Id, ping.NodeType, ping.CurBlock, ping.CurHeight, ping.Time, ping.V, ping.R, ping.S})
}
