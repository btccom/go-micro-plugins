// Package memory provides a memory broker
package memory

import (
	"github.com/btccom/go-micro/v2/broker"
	"github.com/btccom/go-micro/v2/broker/memory"
	"github.com/btccom/go-micro/v2/config/cmd"
)

func init() {
	cmd.DefaultBrokers["memory"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return memory.NewBroker(opts...)
}
