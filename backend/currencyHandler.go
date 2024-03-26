package backend

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/api"
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

func (c *currencyHandler) BalanceOf(ctx *gin.Context) {
	chainId, err := strconv.Atoi(ctx.Param(param_chainId))
	if err != nil {
		ResponseFail(ctx, http.StatusBadRequest, newFailResponse(err, errCode_invalidParams, errMsg_invalidParams))
		return
	}
	currency := common.HexToAddress(ctx.Param(param_currency))
	user := common.HexToAddress(ctx.Param(param_user))

	balance, err := api.GetBalance(context.Background(), c.nodes[chainId], currency, user, nil, c.erc20Abi)
	if err != nil {
		ResponseFail(ctx, http.StatusBadRequest, newFailResponse(err, errCode_networkingNodeFail, errMsg_invalidParams))
		return
	}

	ResponseSuccess(ctx, balance)
}

func (c *currencyHandler) Name(ctx *gin.Context)     {}
func (c *currencyHandler) Symbol(ctx *gin.Context)   {}
func (c *currencyHandler) Transfer(ctx *gin.Context) {}
