module github.com/btccom/go-micro-plugins/broker/segmentio/v2

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/btccom/go-micro/v2 v2.9.3
	github.com/btccom/go-micro-plugins/broker/kafka/v2 v2.3.0
	github.com/btccom/go-micro-plugins/codec/segmentio/v2 v2.3.0
	github.com/segmentio/kafka-go v0.3.7
)

replace github.com/btccom/go-micro-plugins/codec/segmentio/v2 => ../../codec/segmentio
