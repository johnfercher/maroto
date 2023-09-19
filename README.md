# Maroto V2

[![GoDoc](https://godoc.org/github.com/johnfercher/maroto?status.svg)](https://godoc.org/github.com/johnfercher/maroto)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  
[![CI](https://github.com/johnfercher/maroto/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto)
[![Visits Badge](https://badges.pufler.dev/visits/johnfercher/maroto)](https://badges.pufler.dev)

A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

You can write your PDFs like you are creating a site using Bootstrap. A Row may have many Cols, and a Col may have many components. 
Besides that, pages will be added when content may extrapolate the useful area. You can define a header which will be added
always when a new page appear, in this case, a header may have many rows, lines or tablelist. 

* You can see the full documentation [here](https://maroto.io/).
* Discussions are being addressed in [this issue](https://github.com/johnfercher/maroto/issues/257).
* Roadmap is [here](https://github.com/users/johnfercher/projects/1).

![result](docs/assets/images/result.png)

## Installation

* With `go get`:

```bash
go get github.com/johnfercher/maroto/v2/pkg@v2.0.0-alpha.11
```



## Contributing

| Command         | Description                                       | Dependencies                                                  |
|-----------------|---------------------------------------------------|---------------------------------------------------------------|
| `make build`    | Build project                                     | `go`                                                          |
| `make test`     | Run unit tests                                    | `go`                                                          |
| `make fmt`      | Format files                                      | `gofmt`, `gofumpt` and `goimports`                            |
| `make lint`     | Check files                                       | `golangci-lint`                                               |
| `make dod`      | (Definition of Done) Format files and check files | Same as `make build`, `make test`, `make fmt` and `make lint` | 
| `make install`  | Install all dependencies                          | `go`, `curl` and `git`                                        |
| `make font`     | Extract font ut8 to use in development            | `tar`                                                         |
| `make examples` | Run all examples                                  | `go`                                                          |
| `make mock`     | Generate mocks                                    | `go` and `mockery`                                            |
| `make docs`     | Run docs serve local                              | `docsify`                                                     |



## Stargazers over time

[![Stargazers over time](https://starchart.cc/johnfercher/maroto.svg)](https://starchart.cc/johnfercher/maroto)
