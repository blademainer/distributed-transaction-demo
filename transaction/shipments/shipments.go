package shipments

import (
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

const kind = "pay"

// Shipment info
type Shipment struct {
	OrderID string
	Amount  string
	Status  status.Status
}

// Service shipment service
type Service interface {
	// Shipment order
	Shipment(shipment *Shipment) error
}
