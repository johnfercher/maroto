GO_FILES = $(shell find . '(' -path '*/.*' -o -path './vendor' ')' -prune -o -name '*.go' -print | cut -b3-)
GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./... | grep -E -v 'docs|cmd|mocks')

.PHONY: dod
dod: build test fmt lint

.PHONY: build
build:
	go build $(GO_PATHS)

.PHONY: test
test:
	go test $(GO_PATHS)

.PHONY: fmt
fmt:
	gofmt -s -w ${GO_FILES}
	gofumpt -l -w ${GO_FILES}
	goimports -w ${GO_PATHS}

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml ./...

.PHONY: install
install:
	bash install.sh

.PHONY: docs
docs:
	docsify serve docs/

.PHONY: font
font:
	tar -xvf docs/assets/fonts/arial-unicode-ms.tgz

.PHONY: examples
examples: font
	go run docs/assets/examples/barcodegrid/v2/main.go
	# billing
	# billing negative
	go run docs/assets/examples/cellstyle/v2/main.go
	# certificate
	go run docs/assets/examples/customsize/v2/main.go
	go run docs/assets/examples/datamatrixgrid/v2/main.go
	go run docs/assets/examples/footer/v2/main.go
	go run docs/assets/examples/header/v2/main.go
	go run docs/assets/examples/imagegrid/v2/main.go
	go run docs/assets/examples/margins/v2/main.go
	go run docs/assets/examples/maxgridsum/v2/main.go
	go run docs/assets/examples/qrgrid/v2/main.go
	# sample 1
	go run docs/assets/examples/signaturegrid/v2/main.go
	go run docs/assets/examples/textgrid/v2/main.go
	# utf8
	# zpl

.PHONY: mock
mock:
	mockery