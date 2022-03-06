package rpc

// RPC interface
type RPC interface {
	// Invoke remote
	Invoke(req interface{}) error
}
