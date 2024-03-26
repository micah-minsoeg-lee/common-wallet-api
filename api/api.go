package api

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

const (
	method_balanceOf = "balanceOf"
	method_name      = "name"
	method_symbol    = "symbol"
)

func SendTx(ctx context.Context, node *node.Node, signedTx *types.Transaction) (*types.Receipt, error) {
	// send transaction
	if err := node.Client().SendTransaction(ctx, signedTx); err != nil {
		return nil, err
	} else {
		// get receipt
		return node.Client().TransactionReceipt(ctx, signedTx.Hash())
	}
}

func CallContract(ctx context.Context, node *node.Node, callMsg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return node.Client().CallContract(ctx, callMsg, blockNumber)
}

func GetBalance(ctx context.Context, node *node.Node, currency, user common.Address, blockNumber *big.Int, tokenAbi *abi.ABI) (*big.Int, error) {
	if currency == (common.Address{}) {
		// coin
		return node.Client().BalanceAt(ctx, user, blockNumber)
	} else {
		// ERC20 token
		// pack
		if calldata, err := tokenAbi.Pack(method_balanceOf, user); err != nil {
			return nil, fmt.Errorf("pack balanceOf fail: %v", err)
		} else if res, err := CallContract(ctx, node, ethereum.CallMsg{
			To:   &currency,
			From: user,
			Data: calldata,
		}, blockNumber); err != nil {
			return nil, err
		} else {
			return new(big.Int).SetBytes(res), nil
		}
	}
}

func GetName(ctx context.Context, node *node.Node, currency common.Address, tokenAbi *abi.ABI) (string, error) {
	return getCurrencyInfo(ctx, node, method_name, currency, tokenAbi)
}

func GetSymbol(ctx context.Context, node *node.Node, currency common.Address, tokenAbi *abi.ABI) (string, error) {
	return getCurrencyInfo(ctx, node, method_symbol, currency, tokenAbi)
}

func getCurrencyInfo(ctx context.Context, node *node.Node, method string, currency common.Address, tokenAbi *abi.ABI) (string, error) {
	if currency == (common.Address{}) {
		return "", fmt.Errorf("invalid currency address")
	} else {
		// ERC20 token
		// pack
		if calldata, err := tokenAbi.Pack(method); err != nil {
			return "", fmt.Errorf("pack balanceOf fail: %v", err)
		} else if res, err := CallContract(ctx, node, ethereum.CallMsg{
			To:   &currency,
			Data: calldata,
		}, nil); err != nil {
			return "", err
		} else {
			return string(res), nil
		}
	}
}
