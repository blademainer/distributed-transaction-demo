package rpc

import (
	"time"
)

type configuredMockRPC struct {
	TimeAvg         time.Duration
	TimeFluctuation time.Duration
	ErrorPercent    int
}

func (c *configuredMockRPC) Invoke(req interface{}) (interface{}, error) {
	panic("implement me")
}
