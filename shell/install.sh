#!/usr/bin/env bash

# goimports, gofumpt, golangci-lint, mockery and godoc are declared as tool
# dependencies in tools/go.mod and don't need global installation: the
# Makefile invokes them with `go tool -modfile=tools/go.mod <name>`, which
# builds them into the module cache on first use, pinned to the versions
# recorded in tools/go.sum.
#
# This only prefetches them so the first `make fmt`/`make lint`/`make mocks`/
# `make godoc` isn't slowed down by an on-demand build.
go build -modfile=tools/go.mod -o /dev/null \
  golang.org/x/tools/cmd/goimports \
  mvdan.cc/gofumpt \
  github.com/golangci/golangci-lint/v2/cmd/golangci-lint \
  github.com/vektra/mockery/v2 \
  golang.org/x/tools/cmd/godoc

# docsify-cli is a Node.js package (not a Go tool), used to preview docs/.
sudo npm i docsify-cli -g
