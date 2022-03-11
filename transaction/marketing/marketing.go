package marketing

import (
	"github.com/blademainer/distributed-transaction-demo/transaction/status"
)

// Coupon coupon info
type Coupon struct {
	ID             string
	DiscountAmount string
	Status         status.Status
}

type Service interface {
	// PresentCoupon present a coupon to user
	PresentCoupon(userID string, coupon *Coupon) error

	// VerifyCoupon verify a coupon from user
	VerifyCoupon(userID string, coupon *Coupon) error
}
