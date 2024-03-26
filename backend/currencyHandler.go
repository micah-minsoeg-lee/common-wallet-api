package backend

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type currencyHandler struct {
	nodes    node.Nodes
	erc20Abi *abi.ABI
}

func newCurrencyHandler(nodes node.Nodes, tokenAbi *abi.ABI) *currencyHandler {
	return &currencyHandler{
		nodes:    nodes,
		erc20Abi: tokenAbi,
	}
}

func (c *currencyHandler) BalanceOf(ctx *gin.Context) {}
func (c *currencyHandler) Name(ctx *gin.Context)      {}
func (c *currencyHandler) Symbol(ctx *gin.Context)    {}
func (c *currencyHandler) Transfer(ctx *gin.Context)  {}
