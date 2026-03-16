# Low Memory

`WithSequentialLowMemoryMode` switches maroto to a low-memory sequential processing strategy. Instead of holding the entire document in memory while building it, each chunk is processed and flushed to the PDF writer before the next one starts.

## Generation modes comparison

| Mode | Method | Memory | Speed |
|------|--------|--------|-------|
| Default (sequential) | `config.NewBuilder()` | Medium | Baseline |
| Low memory | `WithSequentialLowMemoryMode(n)` | Low | Slightly slower |
| Concurrent | `WithConcurrentMode(workers)` | High | Fastest |

## Usage notes

- Use this mode when generating large documents (hundreds of pages) on memory-constrained environments.
- Headers and footers are still applied to every page; the low-memory mode does not affect their correctness.
- Incompatible with `WithConcurrentMode`; the last mode set wins.

## GoDoc
* [builder : WithSequentialLowMemory](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithSequentialLowMemoryMode)

## Code Example
[filename](../../assets/examples/lowmemory/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/lowmemoryv2.pdf
```

## Time Execution
[filename](../../assets/text/lowmemoryv2.txt  ':include :type=code')


## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/lowmemory.json  ':include :type=code')