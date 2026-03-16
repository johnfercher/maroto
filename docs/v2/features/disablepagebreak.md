# Disable Auto Page Break

By default, maroto automatically inserts a new physical page whenever a row would overflow the current page's usable height. `WithDisableAutoPageBreak` turns off this behaviour so that all content is rendered on a single logical page regardless of how tall it grows.

## Usage notes

- This is useful when generating very long single-scroll documents or when you control pagination entirely through `AddPages`.
- When auto page break is disabled, registered headers and footers are still printed only on the first page.
- Overflowing content is **not** clipped — it extends beyond the bottom margin, which may cause issues in some PDF viewers.
- Re-enable the default behaviour by passing `false` or by omitting the call.

## GoDoc
* [builder : WithDisableAutoPageBreak](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithDisableAutoPageBreak)

## Code Example
[filename](../../assets/examples/disablepagebreak/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/disablepagebreakv2.pdf
```
## Time Execution
[filename](../../assets/text/disablepagebreakv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/disablepagebreak.json  ':include :type=code')