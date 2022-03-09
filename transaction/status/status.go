package status

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
