package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micah-minsoeg-lee/common-wallet-api/backend"
	"github.com/micah-minsoeg-lee/common-wallet-api/config"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
)

type Router struct {
	router   *gin.Engine
	endpoint string
}

func NewRouter(cfg *config.Config, nodes node.Nodes) (*Router, error) {
	// create router
	r := &Router{
		router:   gin.New(),
		endpoint: fmt.Sprintf("%v:%v", cfg.Endpoint.Ip(), cfg.Endpoint.Port()),
	}

	r.router.Use(gin.Logger())   // for using gin logger
	r.router.Use(gin.Recovery()) // for return error code 500 when panic executed

	// TODO: cors

	// create handler
	handler, err := backend.NewHandler(nodes)
	if err != nil {
		return nil, err
	}

	r.bind(handler)

	return r, nil
}

func (r *Router) Router() *gin.Engine {
	return r.router
}

func (r *Router) Endpoint() string {
	return r.endpoint
}

func (r *Router) Run() {
	r.router.Run(r.endpoint)
}

func (r *Router) bind(handler *backend.Handler) {
	chains := r.router.Group("/chains/:chainId")
	{
		block := chains.Group("/block")
		{
			block.GET("/number", handler.BlockHandler().BlockNumber)
		}

		currency := chains.Group("/currency")
		{
			currency.GET("/balance", handler.CurrencyHandler().BalanceOf)
			currency.GET("/name", handler.CurrencyHandler().Name)
			currency.GET("/symbol", handler.CurrencyHandler().Symbol)

			currency.POST("/transfer", handler.CurrencyHandler().Transfer)
		}

		sign := chains.Group("/sign")
		{
			sign.POST("/signTx", handler.SignHandler().SignTx)
			sign.POST("/signMsg", handler.SignHandler().SignMsg)
		}

		tx := chains.Group("/tx")
		{
			tx.POST("/sendTx", handler.TransactionHandler().SendTx)
		}
	}
}
