package pay

// Status order status
type Status int

const (
	Invalid = iota
	Created
	Paying
	Void
	Succeed
	Refunding
	Refunded
)

// Order pay order
type Order struct {
	ID       string
	Amount   string
	Currency string
	Status   Status
}

// Orders orders svc
type Orders interface {
	// Create order
	Create(req *Order) (rsp *Order, err error)

	// Pay update order to success
	Pay(req *Order) (rsp *Order, err error)
}
