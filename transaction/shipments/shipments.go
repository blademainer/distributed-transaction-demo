package shipments

import (
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

// Shipment info
type Shipment struct {
	OrderID string
	Amount  string
	Status  status.Status
}
