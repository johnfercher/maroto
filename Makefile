GO_FILES = $(shell find . '(' -path '*/.*' -o -path './vendor' ')' -prune -o -name '*.go' -print | cut -b3-)
GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./...)

dod: build test fmt lint

build:
	go build -v ./...

test:
	go test -v ./...

fmt:
	gofmt -s -w ${GO_FILES}
	gofumpt -l -w ${GO_FILES}
	goimports -w ${GO_PATHS}

lint:
	goreportcard-cli -v
	golangci-lint run --config=.golangci.yml ./...

install:
	bash install.sh

font:
	tar -xvf docs/assets/fonts/arial-unicode-ms.tgz

v1:
	go run docs/assets/examples/barcodegrid/v1/main.go
	go run docs/assets/examples/imagegrid/v1/main.go
	go run docs/assets/examples/datamatrixgrid/v1/main.go
	go run docs/assets/examples/qrgrid/v1/main.go
	go run docs/assets/examples/textgrid/v1/main.go
	go run docs/assets/examples/maxgridsum/v1/main.go
	go run docs/assets/examples/billing/main.go
	go run docs/assets/examples/billing-with-negative/main.go
	go run docs/assets/examples/certificate/main.go
	go run docs/assets/examples/customsize/main.go
	go run docs/assets/examples/sample1/main.go
	go run docs/assets/examples/zpl/main.go
	go run docs/assets/examples/utfsample/main.go

v2:
	go run docs/assets/examples/barcodegrid/v2/main.go
	go run docs/assets/examples/imagegrid/v2/main.go
	go run docs/assets/examples/datamatrixgrid/v2/main.go
	go run docs/assets/examples/qrgrid/v2/main.go
	go run docs/assets/examples/textgrid/v2/main.go
	go run docs/assets/examples/maxgridsum/v2/main.go