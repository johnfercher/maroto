#!/usr/bin/env bash

go install golang.org/x/tools/cmd/goimports@latest

go install mvdan.cc/gofumpt@latest

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

git clone https://github.com/gojp/goreportcard.git
cd goreportcard
make install
go install ./cmd/goreportcard-cli
cd ..
sudo rm -R goreportcard