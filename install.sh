#!/usr/bin/env bash

go install golang.org/x/tools/cmd/goimports@latest
sudo cp $GOPATH/bin/goimports /usr/local/bin/

go install mvdan.cc/gofumpt@latest
sudo cp $GOPATH/bin/gofumpt /usr/local/bin/

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
sudo cp $GOPATH/bin/golangci-lint /usr/local/bin/

git clone https://github.com/gojp/goreportcard.git
cd goreportcard
make install
go install ./cmd/goreportcard-cli
cd ..
sudo rm -R goreportcard
sudo cp $GOPATH/bin/goreportcard-cli /usr/local/bin/
