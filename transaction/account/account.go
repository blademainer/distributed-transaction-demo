package account

const kind = "account"

// Account info
type Account struct {
	ID    string
	Value string
}

// Service account service
type Service interface {
	// AddValue add value of amount
	AddValue(accountID string, value string) error
}
