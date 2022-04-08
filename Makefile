GO_FILES = $(shell find . '(' -path '*/.*' -o -path './vendor' ')' -prune -o -name '*.go' -print | cut -b3-)
GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./...)

dod: fmt lint

fmt:
	gofmt -s -w ${GO_FILES}
	gofumpt -l -w ${GO_FILES}
	goimports -w ${GO_PATHS}

lint:
	golangci-lint run --config=.golangci.yml ./...
	goreportcard-cli -v

install:
	bash install.sh

examples:
	go run internal/examples/barcodegrid/main.go
	go run internal/examples/billing/main.go
	go run internal/examples/certificate/main.go
	go run internal/examples/customsize/main.go
	go run internal/examples/dmgrid/main.go
	go run internal/examples/imagegrid/main.go
	go run internal/examples/qrgrid/main.go
	go run internal/examples/sample1/main.go
	go run internal/examples/textgrid/main.go
	go run internal/examples/zpl/main.go
	go run internal/examples/utfsample/main.go