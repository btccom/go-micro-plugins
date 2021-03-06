// Package service provides the broker service client
package service

import (
	"github.com/btccom/go-micro/v2/broker"
	"github.com/btccom/go-micro/v2/broker/service"
	"github.com/btccom/go-micro/v2/config/cmd"
)

func init() {
	cmd.DefaultBrokers["service"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return service.NewBroker(opts...)
}
