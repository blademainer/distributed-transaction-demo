package marketing

import (
	"github.com/blademainer/distributed-transaction-demo/mock/rpc"
)

type mocked struct {
	rpc rpc.RPC
}

func (m *mocked) PresentCoupon(userID string, coupon *Coupon) error {
	_, err := m.rpc.Invoke([]interface{}{userID, coupon})
	if err != nil {
		return err
	}

	return nil
}

func (m *mocked) VerifyCoupon(userID string, coupon *Coupon) error {
	_, err := m.rpc.Invoke([]interface{}{userID, coupon})
	if err != nil {
		return err
	}

	return nil
}

// NewMarketingServer create marketing service
func NewMarketingServer(rpc rpc.RPC) Service {
	orders := &mocked{
		rpc: rpc,
	}
	return orders
}
