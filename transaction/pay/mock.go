package pay

import (
	"github.com/blademainer/distributed-transaction-demo/mock/rpc"
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

type mockedOrders struct {
	rpc rpc.RPC
}

// NewOrdersServer create orders service
func NewOrdersServer(rpc rpc.RPC) Service {
	orders := &mockedOrders{
		rpc: rpc,
	}
	return orders
}

// Create new order
func (m *mockedOrders) Create(req *Order) (rsp *Order, err error) {
	_, err = m.rpc.Invoke(req)
	if err != nil {
		return nil, err
	}
	rsp = &Order{
		ID:       req.ID,
		Amount:   req.Amount,
		Currency: req.Currency,
		Status:   status.Created,
	}
	return
}

func (m *mockedOrders) Pay(req *Order) (rsp *Order, err error) {
	_, err = m.rpc.Invoke(req)
	if err != nil {
		return nil, err
	}
	rsp = &Order{
		ID:       req.ID,
		Amount:   req.Amount,
		Currency: req.Currency,
		Status:   status.Succeed,
	}
	return
}
