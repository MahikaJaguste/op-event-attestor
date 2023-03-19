package derive

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	CustomEventABI       = "Transfer(address,address,uint256)"
	CustomEventABIHash   = crypto.Keccak256Hash([]byte(CustomEventABI))
	CustomEventVersion0  = common.Hash{}
)
