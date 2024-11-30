package eth

import (
	"context"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

type PublicSafe3API struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSafe3API(e *Ethereum) *PublicSafe3API {
	return &PublicSafe3API{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func Sign(safe3Addr string, privkey []byte) []byte {
	h := sha256.Sum256([]byte(safe3Addr))
	h2 := accounts.TextHash(h[:])
	sig, _ := secp256k1.Sign(h2, privkey)
	sig[crypto.RecoveryIDOffset] += 27
	return sig
}

func parseKey(key string) []byte {
	//fmt.Printf("%s\n", hexutils.BytesToHex(base58.Decode(key)))
	//fmt.Printf("%s\n", hexutils.BytesToHex(base58.Decode(key)[1:33]))
	return base58.Decode(key)[1:33]
}

func (api *PublicSafe3API) RedeemWithKeys(ctx context.Context, from common.Address, keys []string, targetAddr common.Address) ([]common.Hash, error) {
	availablePubkeys := make([]hexutil.Bytes, 0)
	availableSigs := make([]hexutil.Bytes, 0)
	lockedPubkeys := make([]hexutil.Bytes, 0)
	lockedSigs := make([]hexutil.Bytes, 0)
	mnPubkeys := make([]hexutil.Bytes, 0)
	mnSigs := make([]hexutil.Bytes, 0)
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	for _, key := range keys {
		privkey := parseKey(key)
		priv, _ := crypto.ToECDSA(privkey)

		// compressed pubkey
		pubkey := crypto.CompressPubkey(&priv.PublicKey)
		safe3Addr := getSafe3Addr(pubkey)
		flag, err := api.ExistAvailableNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			availablePubkeys = append(availablePubkeys, pubkey)
			availableSigs = append(availableSigs, Sign(safe3Addr, privkey))
		}
		flag, err = api.ExistLockedNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			lockedPubkeys = append(lockedPubkeys, pubkey)
			lockedSigs = append(lockedSigs, Sign(safe3Addr, privkey))
		}
		flag, err = api.ExistMasterNodeNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			mnPubkeys = append(mnPubkeys, pubkey)
			mnSigs = append(mnSigs, Sign(safe3Addr, privkey))
		}

		// uncompressed pubkey
		pubkey = crypto.FromECDSAPub(&priv.PublicKey)
		safe3Addr = getSafe3Addr(pubkey)
		flag, err = api.ExistAvailableNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			availablePubkeys = append(availablePubkeys, pubkey)
			availableSigs = append(availableSigs, Sign(safe3Addr, privkey))
		}
		flag, err = api.ExistLockedNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			lockedPubkeys = append(lockedPubkeys, pubkey)
			lockedSigs = append(lockedSigs, Sign(safe3Addr, privkey))
		}
		flag, err = api.ExistMasterNodeNeedToRedeem(ctx, safe3Addr, blockNrOrHash)
		if err == nil && flag {
			mnPubkeys = append(mnPubkeys, pubkey)
			mnSigs = append(mnSigs, Sign(safe3Addr, privkey))
		}
	}

	txs := make([]common.Hash, 0)
	if len(availablePubkeys) > 0 {
		if availableTx, err := contract_api.BatchRedeemAvailable(ctx, api.blockChainAPI, api.transactionPoolAPI, from, availablePubkeys, availableSigs, targetAddr); err == nil {
			txs = append(txs, availableTx)
		}
	}
	if len(lockedPubkeys) > 0 {
		if lockedTx, err := contract_api.BatchRedeemLocked(ctx, api.blockChainAPI, api.transactionPoolAPI, from, lockedPubkeys, lockedSigs, targetAddr); err == nil {
			txs = append(txs, lockedTx)
		}
	}
	if len(mnPubkeys) > 0 {
		mnEnodes := make([]string, len(mnPubkeys))
		for i, _ := range mnPubkeys {
			mnEnodes[i] = ""
		}
		if mnTx, err := contract_api.BatchRedeemMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, lockedPubkeys, lockedSigs, mnEnodes, targetAddr); err == nil {
			txs = append(txs, mnTx)
		}
	}
	return txs, nil
}

func (api *PublicSafe3API) ApplyRedeemSpecialWithKey(ctx context.Context, from common.Address, key string) ([]common.Hash, error) {
	privkey := parseKey(key)
	priv, _ := crypto.ToECDSA(privkey)

	txs := make([]common.Hash, 0)
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)

	// compressed pubkey
	pubkey := crypto.CompressPubkey(&priv.PublicKey)
	safe3Addr := getSafe3Addr(pubkey)
	specialInfo, err := contract_api.GetSpecialInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
	if err == nil && specialInfo.Amount.Int64() != 0 && specialInfo.RedeemHeight.Int64() == 0 {
		txid, _ := contract_api.ApplyRedeemSpecial(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, Sign(safe3Addr, privkey))
		if txid != (common.Hash{}) {
			txs = append(txs, txid)
		}
	}

	// uncompressed pubkey
	pubkey = crypto.FromECDSAPub(&priv.PublicKey)
	safe3Addr = getSafe3Addr(pubkey)
	specialInfo, err = contract_api.GetSpecialInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
	if err == nil && specialInfo.Amount.Int64() != 0 && specialInfo.RedeemHeight.Int64() == 0 {
		txid, _ := contract_api.ApplyRedeemSpecial(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, Sign(safe3Addr, privkey))
		if txid != (common.Hash{}) {
			txs = append(txs, txid)
		}
	}
	return txs, nil
}

func (api *PublicSafe3API) Vote4Special(ctx context.Context, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4Special(ctx, api.blockChainAPI, api.transactionPoolAPI, from, safe3Addr, voteResult)
}

func (api *PublicSafe3API) GetAllAvailableNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllAvailableNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetAvailableInfos(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
	return contract_api.GetAvailableInfos(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetAvailableInfo(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
	return contract_api.GetAvailableInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllLockedNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllLockedNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedAddrNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetLockedAddrNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedAddrs(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	return contract_api.GetLockedAddrs(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedNum(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetLockedNum(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedInfo(ctx context.Context, safe3Addr string, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
	return contract_api.GetLockedInfo(ctx, api.blockChainAPI, safe3Addr, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllSpecialNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllSpecialNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetSpecialInfos(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
	return contract_api.GetSpecialInfos(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetSpecialInfo(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
	return contract_api.GetSpecialInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) ExistAvailableNeedToRedeem(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistAvailableNeedToRedeem(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) ExistLockedNeedToRedeem(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistLockedNeedToRedeem(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) ExistMasterNodeNeedToRedeem(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistMasterNodeNeedToRedeem(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func getKeyID(pubkey []byte) []byte {
	//fmt.Printf("%s\n", hexutils.BytesToHex(pubkey))
	h := sha256.Sum256(pubkey)
	//fmt.Printf("hash: %s\n", hexutils.BytesToHex(h[:]))
	ripemd := ripemd160.New()
	ripemd.Write(h[:])
	r := ripemd.Sum(nil)
	//fmt.Printf("ripemd160: %s\n", hexutils.BytesToHex(r))
	t := append([]byte{0x4c}, r...)
	h = sha256.Sum256(t)
	h = sha256.Sum256(h[:])
	//fmt.Printf("h: %s\n", hexutils.BytesToHex(h[:]))
	t = append(t, h[0:4]...)
	//fmt.Printf("t: %s\n", hexutils.BytesToHex(t))
	return t
}

func getSafe3Addr(pubkey []byte) string {
	return base58.Encode(getKeyID(pubkey))
}

func getSafe4Addr(privkey []byte) common.Address {
	priv, _ := crypto.ToECDSA(secp256k1.LoadKey(privkey))
	return crypto.PubkeyToAddress(priv.PublicKey)
}
