package rpc

// RPC interface
type RPC interface {
	// Invoke remote
	Invoke(kind string, key string, req interface{}) (rsp interface{}, err error)
}
