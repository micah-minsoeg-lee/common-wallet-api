package backend

import "github.com/micah-minsoeg-lee/common-wallet-api/node"

type Handler struct {
	currencyHandler    *currencyHandler
	transactionHandler *transactionHandler
	signHandler        *signHandler
	blockHandler       *blockHandler
}

func NewHandler(nodes node.Nodes) (*Handler, error) {
	return &Handler{
		currencyHandler:    newCurrencyHandler(nodes),
		transactionHandler: newTransactionHandler(nodes),
		signHandler:        newSignHandler(nodes),
		blockHandler:       newBlockHandler(nodes),
	}, nil
}

func (h *Handler) CurrencyHandler() *currencyHandler {
	return h.currencyHandler
}

func (h *Handler) TransactionHandler() *transactionHandler {
	return h.transactionHandler
}

func (h *Handler) SignHandler() *signHandler {
	return h.signHandler
}

func (h *Handler) BlockHandler() *blockHandler {
	return h.blockHandler
}
