package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type nodeInfo struct {
	url string
}

func (n *nodeInfo) Url() string {
	return n.url
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
	Nodes map[string]nodeInfo // chain name => chain node

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
		c.Endpoint.port = 8080
	}
}
