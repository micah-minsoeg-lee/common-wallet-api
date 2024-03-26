package backend

import (
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
