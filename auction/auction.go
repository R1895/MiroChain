package auction

import (
	"MicroChain/credit"
	"fmt"
)

func DoubleAuction() {
	credit.AdjustCredit()
	fmt.Println("DoubleAuction")
}
