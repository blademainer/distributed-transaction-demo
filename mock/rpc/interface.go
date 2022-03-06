package rpc

// RPC interface
type RPC interface {
	// Invoke remote
	Invoke(req interface{}) (rsp interface{}, err error)
}
