#!/bin/bash -e

tag=$1

for m in $(find ./ -name 'go.mod'); do
  d=$(dirname $m);
  pushd $d;
  grep -q github.com/btccom/go-micro/v2 go.mod && go get github.com/btccom/go-micro/v2@$tag
  grep -q github.com/btccom/go-micro-platform/v2 go.mod && go get github.com/btccom/go-micro-platform/v2@$tag
  go mod tidy
  popd;
done
