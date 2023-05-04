#!/bin/sh

mkdir -p github.com/gogo
pushd github.com/gogo
git clone https://github.com/gogo/protobuf
popd

~/go/bin/go-to-protobuf -h boilerplate.txt -p github.com/erhudy/goboolstr --apimachinery-packages -

mv github.com/erhudy/goboolstr/generated.pb.go github.com/erhudy/goboolstr/generated.proto .

rm -rf github.com