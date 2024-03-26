package backend

import (
	"fmt"
	"math/big"
)

// def params name
const (
	param_chainId  = "chainId"
	param_currency = "currencyAddress"
	param_user     = "userAddress"
)

// def error response code
const (
	// TODO: def custom err code
	errCode_invalidParams      = -1000
	errCode_invalidRequest     = -1001
	errCode_networkingNodeFail = -2000
)

// def error response msg
const (
	errMsg_invalidParams      = "invalid params"
	errMsg_invalidRequest     = "invalid request"
	errMsg_networkingNodeFail = "networking node fail"

	errMsg_invalidFeeCap = "invalid feecap"
	errMsg_invalidTipCap = "invalid tipcap"
)

type failResponse struct {
	errCode int
	message string
}

func newFailResponse(err error, errCode int, msg string) failResponse {
	return failResponse{
		errCode: errCode,
		message: fmt.Sprintf("%v: %v", msg, err),
	}
}

type currencyTransferRequest struct {
	Currency string `json:"currency,omitempty"`
	To       string `json:"to"`
	Value    string `json:"value"`
}

type gasPriceInfoSet struct {
	FeeCap string `json:"feecap,maxFeePerGas,gasPrice"`
	TipCap string `json:"tipcap,maxPriorityFeePerGas,omitempty"`
}

type gasPriceInfo struct {
	feeCap *big.Int
	tipCap *big.Int
}

func convertInfoSetToInfo(set gasPriceInfoSet) (gasPriceInfo, error) {
	if set.TipCap == "" {
		set.TipCap = set.FeeCap
	}

	feeCap, ok := new(big.Int).SetString(set.FeeCap, 16)
	if !ok {
		return gasPriceInfo{}, fmt.Errorf(errMsg_invalidFeeCap)
	}

	tipCap, ok := new(big.Int).SetString(set.TipCap, 16)
	if !ok {
		return gasPriceInfo{}, fmt.Errorf(errMsg_invalidTipCap)
	}

	return gasPriceInfo{
		feeCap: feeCap,
		tipCap: tipCap,
	}, nil
}

type msgSignPrefix string

const (
	eip_712_ethereum_prefix  msgSignPrefix = "\x19Ethereum Signed Message:\n"
	eip_712_avalanche_prefix msgSignPrefix = "\x1AAvalanche Signed Message:\n"
)

func prefix(name string) msgSignPrefix {
	if name == "avalanche" {
		return eip_712_avalanche_prefix
	} else {
		return eip_712_ethereum_prefix
	}
}
