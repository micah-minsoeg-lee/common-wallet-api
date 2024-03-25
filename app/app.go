package app

import (
	"github.com/micah-minsoeg-lee/common-wallet-api/app/flag"

	"github.com/micah-minsoeg-lee/common-wallet-api/config"
	"github.com/micah-minsoeg-lee/common-wallet-api/node"
	"github.com/micah-minsoeg-lee/common-wallet-api/router"
)

type App struct {
	router *router.Router
	config *config.Config
	nodes  node.Nodes
}

func NewApp() (*App, error) {
	// create app
	app := new(App)

	// parse flag
	flags := flag.Parse()

	var err error
	// create config obj
	if app.config, err = config.NewConfig(flags.ConfigFlag()); err != nil {
		return nil, err
	}

	// create nodes
	if app.nodes, err = node.NewNodes(app.config); err != nil {
		return nil, err
	}

	// create router
	if app.router, err = router.NewRouter(app.config, app.nodes); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() {
	a.router.Run()
}
