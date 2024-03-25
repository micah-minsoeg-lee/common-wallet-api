package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type currencyHandler struct {
	nodes node.Nodes
}

func newCurrencyHandler(nodes node.Nodes) *currencyHandler {
	return &currencyHandler{
		nodes: nodes,
	}
}

func (c *currencyHandler) BalanceOf(ctx *gin.Context) {}
func (c *currencyHandler) Name(ctx *gin.Context)      {}
func (c *currencyHandler) Symbol(ctx *gin.Context)    {}
