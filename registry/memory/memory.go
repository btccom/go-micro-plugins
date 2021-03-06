// Package memory provides an in-memory registry
package memory

import (
	"github.com/btccom/go-micro/v2/config/cmd"
	"github.com/btccom/go-micro/v2/registry"
	"github.com/btccom/go-micro/v2/registry/memory"
)

func init() {
	cmd.DefaultRegistries["memory"] = NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return memory.NewRegistry(opts...)
}
