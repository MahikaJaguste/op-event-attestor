package derive

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)


func DeriveCustomEvents(receipts []*types.Receipt) () {
	for _, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for _, log := range rec.Logs {
			if len(log.Topics) > 0 && log.Topics[0] == CustomEventABIHash {
				fmt.Println("mahika log", log.Topics)
			}
		}
	}
}
