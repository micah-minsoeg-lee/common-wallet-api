package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type transactionHandler struct {
	nodes node.Nodes
}

func newTransactionHandler(nodes node.Nodes) *transactionHandler {
	return &transactionHandler{
		nodes: nodes,
	}
}

func (t *transactionHandler) SendTx(ctx *gin.Context) {}
