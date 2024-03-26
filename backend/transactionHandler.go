package backend

import (
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
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

func (t *transactionHandler) MakeUnsignedTx(ctx *gin.Context) {
	req := new(currencyTransferRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseFail(ctx, http.StatusBadRequest, newFailResponse(err, errCode_invalidRequest, errMsg_invalidRequest))
		return
	}

	ResponseSuccess(ctx, types.NewTx(&types.DynamicFeeTx{}))
}

func (t *transactionHandler) SendTx(ctx *gin.Context) {

}
