package abi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const erc20AbiPath = "./ERC20.json"

func GetErc20Abi() (*abi.ABI, error) {
	// get erc20 token abi
	tokenAbi := new(abi.ABI)
	if file, err := os.Open(path.Join(erc20AbiPath)); err != nil {
		return nil, fmt.Errorf("open erc20 token abi fail: %v", err)
	} else {
		defer file.Close()
		// read file
		abiBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("read erc20 token abi fail: %v", err)
		}
		// unmarshal abi
		err = json.Unmarshal(abiBytes, tokenAbi)
		if err != nil {
			return nil, fmt.Errorf("unmarshal erc20 token abi fail: %v", err)
		}
	}
	return tokenAbi, nil
}
