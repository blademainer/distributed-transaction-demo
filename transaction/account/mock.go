package account

import (
	"sync"

	"github.com/blademainer/distributed-transaction-demo/mock/rpc"
)

type mocked struct {
	rpc      rpc.RPC
	accounts *sync.Map
}

// NewAccountServer create Account service
func NewAccountServer(rpc rpc.RPC) Service {
	orders := &mocked{
		rpc:      rpc,
		accounts: &sync.Map{},
	}
	return orders
}

func (m *mocked) AddValue(accountID string, value string) error {
	_, err := m.rpc.Invoke(kind, accountID, value)
	if err != nil {
		return err
	}

	return nil
}
