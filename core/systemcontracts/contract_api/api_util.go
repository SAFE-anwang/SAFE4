package contract_api

import (
	"context"
    "math/big"

    "github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetCurrentGasPrice(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI) *big.Int {
	gasPrice, err := GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || gasPrice.Int64() == 0 {
		gasPrice = big.NewInt(params.GWei / 100)
	}
	return gasPrice
}
