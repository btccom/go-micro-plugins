package etcdv3

import (
	"github.com/btccom/go-micro/v2/registry"
	"github.com/btccom/go-micro/v2/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
