package node

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/micah-minsoeg-lee/common-wallet-api/config"
)

type Node struct {
	client *ethclient.Client
	info   *config.Currency
}

func NewNode(path string, info *config.Currency) (*Node, error) {
	n := new(Node)
	var err error
	if n.client, err = ethclient.DialContext(context.Background(), path); err != nil {
		return nil, err
	}

	return n, nil
}

func (n *Node) Client() *ethclient.Client {
	return n.client
}

func (n *Node) Info() *config.Currency {
	return n.info
}

type Nodes map[int]*Node

func NewNodes(cfg *config.Config) (Nodes, error) {
	nodes := make(Nodes)

	for key, info := range cfg.Nodes {
		node, err := NewNode(info.Url(), info.CurrencyInfo())
		if err != nil {
			return nil, err
		}
		nodes[key] = node
	}

	return nodes, nil
}
