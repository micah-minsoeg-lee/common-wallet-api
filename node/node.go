package node

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/micah-minsoeg-lee/common-wallet-api/config"
)

type Node struct {
	client *ethclient.Client
}

func NewNode(path string) (*Node, error) {
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

type Nodes map[int]*Node

func NewNodes(cfg *config.Config) (Nodes, error) {
	nodes := make(Nodes)

	for key, info := range cfg.Nodes {
		node, err := NewNode(info.Url())
		if err != nil {
			return nil, err
		}
		nodes[key] = node
	}

	return nodes, nil
}
