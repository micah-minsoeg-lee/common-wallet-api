package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type blockHandler struct {
	nodes node.Nodes
}

func newBlockHandler(nodes node.Nodes) *blockHandler {
	return &blockHandler{
		nodes: nodes,
	}
}

func (b *blockHandler) BlockNumber(ctx *gin.Context) {}
