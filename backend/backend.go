package backend

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/micah-minsoeg-lee/common-wallet-api/abi"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type Handler struct {
	currencyHandler    *currencyHandler
	transactionHandler *transactionHandler
	blockHandler       *blockHandler
}

func NewHandler(nodes node.Nodes) (*Handler, error) {
	// get erc20 token abi
	tokenAbi, err := abi.GetErc20Abi()
	if err != nil {
		return nil, err
	}
	return &Handler{
		currencyHandler:    newCurrencyHandler(nodes, tokenAbi),
		transactionHandler: newTransactionHandler(nodes),
		blockHandler:       newBlockHandler(nodes),
	}, nil
}

func (h *Handler) CurrencyHandler() *currencyHandler {
	return h.currencyHandler
}

func (h *Handler) TransactionHandler() *transactionHandler {
	return h.transactionHandler
}

func (h *Handler) BlockHandler() *blockHandler {
	return h.blockHandler
}

func makeUnsignedTx(to common.Address, chainId, value *big.Int, nonce, gas uint64, gasInfo gasPriceInfo, data []byte) *types.Transaction {
	// TODO: upgrade to more details
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasInfo.tipCap,
		GasFeeCap: gasInfo.feeCap,
		Gas:       gas,
		Value:     value,
		Data:      data,
	})
}
