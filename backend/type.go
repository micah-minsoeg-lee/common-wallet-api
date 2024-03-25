package backend

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
