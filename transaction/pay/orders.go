package pay

import (
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

const kind = "pay"

// Order pay order
type Order struct {
	ID       string
	Amount   string
	Currency string
	Status   status.Status
}

// Service orders svc
type Service interface {
	// Create order
	Create(req *Order) (rsp *Order, err error)

	// Pay update order to success
	Pay(req *Order) (rsp *Order, err error)
}
