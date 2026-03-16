# Parallelism

`WithConcurrentMode` enables concurrent PDF generation. maroto splits the document into chunks and processes them in parallel using a configurable number of worker goroutines, then assembles the final output in the correct order.

## Generation modes comparison

| Mode | Method | Memory | Speed |
|------|--------|--------|-------|
| Default (sequential) | `config.NewBuilder()` | Medium | Baseline |
| Low memory | `WithSequentialLowMemoryMode(n)` | Low | Slightly slower |
| Concurrent | `WithConcurrentMode(workers)` | High | Fastest |

## Usage notes

- Concurrent mode provides the most significant speed gains for large documents (50+ pages) or documents with heavy per-row computation.
- Memory usage scales with the number of workers because each worker holds its chunk in memory simultaneously.
- Incompatible with `WithSequentialLowMemoryMode`; the last one called wins.

## GoDoc
* [builder : WithConcurrentMode](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithConcurrentMode)

## Code Example
[filename](../../assets/examples/parallelism/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/parallelismv2.pdf
```

## Time Execution
[filename](../../assets/text/parallelismv2.txt  ':include :type=code')

## Time Execution (100 pages)
[filename](../../assets/text/parallel.txt ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/parallelism.json  ':include :type=code')