package backend

import "fmt"

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
	errCode_networkingNodeFail = -2000
)

// def error response msg
const (
	errMsg_invalidParams      = "invalid params"
	errMsg_networkingNodeFail = "networking node fail"
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
