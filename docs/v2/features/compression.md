# Compression

`WithCompression` enables zlib compression on the generated PDF. Compression reduces the output file size — particularly effective for documents with repeated patterns, long text, or many identical glyphs — at the cost of slightly more CPU time during generation.

Compression is **disabled** by default. Pass `true` to enable it or `false` to explicitly disable it.

## Usage notes

- Compression is applied to the entire PDF stream; individual image compression is independent of this setting.
- For documents that are predominantly images, the size reduction may be modest because images are usually already compressed.
- Enable compression whenever the output PDF will be stored or transmitted and file size matters.

## GoDoc
* [builder : WithCompression](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithCompression)

## Code Example
[filename](../../assets/examples/compression/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/compressionv2.pdf
```
## Time Execution
[filename](../../assets/text/compressionv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/compression.json  ':include :type=code')