// Package memory provides a memory broker
package memory

import (
	"github.com/btccom/go-micro/v2/broker"
	"github.com/btccom/go-micro/v2/broker/nats"
	"github.com/btccom/go-micro/v2/config/cmd"
)

func init() {
	cmd.DefaultBrokers["nats"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return nats.NewBroker(opts...)
}
