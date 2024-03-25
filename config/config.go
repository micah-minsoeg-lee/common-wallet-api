package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type currency struct {
	name   string
	symbol string
}

func (c *currency) Name() string {
	return c.name
}

func (c *currency) Symbol() string {
	return c.symbol
}

type nodeInfo struct {
	url          string
	name         string
	currencyInfo currency
}

func (n *nodeInfo) Url() string {
	return n.url
}

func (n *nodeInfo) Name() string {
	return n.name
}

func (n *nodeInfo) CurrencyInfo() currency {
	return n.currencyInfo
}

type endpoint struct {
	ip   string
	port int
}

func (e *endpoint) Ip() string {
	return e.ip
}

func (e *endpoint) Port() int {
	return e.port
}

const defaultPath = "./config.toml"

type Config struct {
	Nodes map[int]nodeInfo // chain id => chain node info

	Endpoint endpoint
}

func NewConfig(filePath string) (*Config, error) {
	if filePath == "" {
		filePath = defaultPath
	}

	if file, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer file.Close()

		c := new(Config)
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			return nil, err
		} else {
			c.setDefaultValue()
			return c, err
		}
	}
}

func (c *Config) setDefaultValue() {
	if c.Endpoint.ip == "" {
		// localhost
		c.Endpoint.ip = "127.0.0.1"
	} else if c.Endpoint.port == 0 {
		// default port
		c.Endpoint.port = 8080
	}
}
