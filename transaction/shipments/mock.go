package shipments

import (
	"sync"

	"github.com/blademainer/distributed-transaction-demo/mock/rpc"
)

type mocked struct {
	rpc      rpc.RPC
	accounts *sync.Map
}

func (m *mocked) Shipment(shipment *Shipment) error {
	_, err := m.rpc.Invoke(shipment)
	if err != nil {
		return err
	}

	return nil
}

// NewShipmentsServer create orders service
func NewShipmentsServer(rpc rpc.RPC) Service {
	shipments := &mocked{
		rpc:      rpc,
		accounts: &sync.Map{},
	}
	return shipments
}
