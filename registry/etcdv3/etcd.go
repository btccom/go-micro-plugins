// Package etcd provides an etcd v3 service registry
package etcdv3

import (
	"github.com/btccom/go-micro/v2/config/cmd"
	"github.com/btccom/go-micro/v2/registry"
	"github.com/btccom/go-micro/v2/registry/etcd"
)

func init() {
	cmd.DefaultRegistries["etcdv3"] = etcd.NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return etcd.NewRegistry(opts...)
}
