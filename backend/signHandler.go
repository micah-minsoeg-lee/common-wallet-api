package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type signHandler struct {
	nodes node.Nodes
}

func newSignHandler(nodes node.Nodes) *signHandler {
	return &signHandler{
		nodes: nodes,
	}
}

func (s *signHandler) SignTx(ctx *gin.Context) {

}

func (s *signHandler) SignMsg(ctx *gin.Context) {

}
