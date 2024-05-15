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
	make mock-lint

.PHONY: mock-lint
mock-lint:
	bash shell/mock-check.sh

.PHONY: install
install:
	bash shell/install.sh

.PHONY: docs
docs:
	docsify serve docs/

.PHONY: godoc
godoc:
	godoc -http=127.0.0.1:6060


.PHONY: mocks
mocks:
	mockery

.PHONY: examples
examples:
	go run docs/assets/examples/addpage/v2/main.go
	go run docs/assets/examples/background/v2/main.go
	go run docs/assets/examples/barcodegrid/v2/main.go
	go run docs/assets/examples/billing/v2/main.go
	go run docs/assets/examples/cellstyle/v2/main.go
	go run docs/assets/examples/compression/v2/main.go
	go run docs/assets/examples/customdimensions/v2/main.go
	go run docs/assets/examples/customfont/v2/main.go
	go run docs/assets/examples/custompage/v2/main.go
	go run docs/assets/examples/datamatrixgrid/v2/main.go
	go run docs/assets/examples/disablepagebreak/v2/main.go
	go run docs/assets/examples/footer/v2/main.go
	go run docs/assets/examples/header/v2/main.go
	go run docs/assets/examples/imagegrid/v2/main.go
	go run docs/assets/examples/line/v2/main.go
	go run docs/assets/examples/list/v2/main.go
	go run docs/assets/examples/margins/v2/main.go
	go run docs/assets/examples/maxgridsum/v2/main.go
	go run docs/assets/examples/mergepdf/v2/main.go
	go run docs/assets/examples/metadatas/v2/main.go
	go run docs/assets/examples/orientation/v2/main.go
	go run docs/assets/examples/pagenumber/v2/main.go
	go run docs/assets/examples/parallelism/v2/main.go
	go run docs/assets/examples/protection/v2/main.go
	go run docs/assets/examples/qrgrid/v2/main.go
	go run docs/assets/examples/signaturegrid/v2/main.go
	go run docs/assets/examples/simplest/v2/main.go
	go run docs/assets/examples/textgrid/v2/main.go
	go test docs/assets/examples/unittests/v2/main_test.go
