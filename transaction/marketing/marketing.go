package marketing

import (
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

type Coupons struct {
	ID             string
	DiscountAmount string
	Status         status.Status
}
