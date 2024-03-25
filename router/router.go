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
	handler  *backend.Handler
	endpoint string
}

func NewRouter(cfg *config.Config, nodes node.Nodes) (*Router, error) {
	// create router
	r := &Router{
		router:   gin.New(),
		endpoint: fmt.Sprintf("%v:%v", cfg.Endpoint.Ip(), cfg.Endpoint.Port()),
	}

	var err error
	// create handler
	if r.handler, err = backend.NewHandler(nodes); err != nil {
		return nil, err
	}

	r.router.Use(gin.Logger())   // for using gin logger
	r.router.Use(gin.Recovery()) // for return error code 500 when panic executed

	// TODO: cors

	return r, nil
}

func (r *Router) Router() *gin.Engine {
	return r.router
}

func (r *Router) Handler() *backend.Handler {
	return r.handler
}

func (r *Router) Endpoint() string {
	return r.endpoint
}

func (r *Router) Run() {
	r.router.Run(r.endpoint)
}
