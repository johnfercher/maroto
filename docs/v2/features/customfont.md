# Custom Font

Maroto ships with a standard set of built-in fonts (`arial`, `helvetica`, `courier`, `times`, `symbol` and `zapfdingbats`), rendered with the embedded Liberation fonts, which have the same metrics as Arial/Helvetica, Courier and Times and cover Latin, Greek and Cyrillic. `WithCustomFonts` lets you register additional TrueType (UTF-8) fonts so that text components can use them by name. This is essential for branded typography and for scripts the built-in fonts do not cover, such as Arabic, Chinese, Japanese or Korean.

## Usage notes

- Each style variant (`Normal`, `Bold`, `Italic`, `BoldItalic`) must be registered separately.
- If a style is not registered, `Generate` fails: with an `undefined font` error in the default sequential mode, with the mode's own error (`ErrCannotGenerateInParallelMode`, `ErrCannotGenerateInLowMemoryMode`) in the concurrent and low-memory modes. A custom font registered with the name of a built-in family (e.g. `arial`) replaces it, for the registered styles only.
- Only the glyphs used by the document are embedded in the PDF (font subsetting), so the output is self-contained.
- Each family and style is parsed once per process and shared by every document that registers the same bytes for it, so generating many documents with large fonts does not parse them again. A file registered for several styles is kept in memory once.
- Only TrueType outlines (`.ttf`) are supported; OpenType fonts with CFF outlines (`.otf`) are rejected.

## GoDoc
* [builder : WithCustomFonts](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithCustomFonts)
* [repository : AddUTF8Font](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/repository#FontRepository.AddUTF8Font)
* [repository : Load](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/repository#FontRepository.Load)
* [entity : CustomFont](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core/entity#CustomFont)

## Code Example
[filename](../../assets/examples/customfont/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/customfontv2.pdf
```
## Time Execution
[filename](../../assets/text/customfontv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/customfont.json  ':include :type=code')