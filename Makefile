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

v1:
	go run internal/examples/barcodegrid/v1/main.go
	go run internal/examples/imagegrid/v1/main.go
	go run internal/examples/datamatrixgrid/v1/main.go
	go run internal/examples/qrgrid/v1/main.go
	go run internal/examples/textgrid/v1/main.go
	go run internal/examples/billing/main.go
	go run internal/examples/billing-with-negative/main.go
	go run internal/examples/certificate/main.go
	go run internal/examples/customsize/main.go
	go run internal/examples/sample1/main.go
	go run internal/examples/zpl/main.go
	go run internal/examples/maxgridsum/main.go
	go run internal/examples/utfsample/main.go

v2:
	go run internal/examples/barcodegrid/v2/main.go
	go run internal/examples/imagegrid/v2/main.go
	go run internal/examples/datamatrixgrid/v2/main.go
	go run internal/examples/qrgrid/v2/main.go
	go run internal/examples/textgrid/v2/main.go