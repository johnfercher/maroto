# Arabic (RTL)

Arabic is written from right to left, and each letter changes shape depending on
the letters it connects to. PDF text operators know neither: they draw glyphs one
after another from left to right, with no script specific logic. Arabic passed
straight through therefore comes out as disconnected letters in reversed order.

Setting `RTL` on `props.Text` turns on the two transformations that fix it:

* **shaping**, which replaces every letter by the contextual presentation form
  that connects to its neighbours and contracts the mandatory lam-alef ligatures;
* **reordering**, which lays the bidirectional runs out in the visual order the
  writer has to emit.

```golang
m.AddRows(text.NewRow(10, "مرحبا بالعالم", props.Text{
    Align: align.Right,
    RTL:   true,
}))
```

The same flag exists on `props.Checkbox` for checkbox labels.

## Fonts

The standard PDF fonts do not carry Arabic glyphs, so a custom font that covers
the Arabic Presentation Forms-B block (U+FE70 to U+FEFF) is required. See
[Custom Font](v2/features/customfont.md?id=custom-font) for how to register one.

```golang
customFonts, err := fontrepository.New().
    AddUTF8Font("arial-unicode-ms", fontstyle.Normal, "font.ttf").
    Load()
```

## Behaviour

| Situation | Result |
|-----------|--------|
| `RTL` not set | The text is written unchanged, exactly as before |
| `RTL` set, text has no Arabic | The text is returned byte for byte unchanged |
| `RTL` set, text has Arabic | The text is shaped and reordered |
| Text already shaped by the caller | Do not set `RTL`, it would be processed twice |

The flag is opt-in for the last reason: applications that already feed maroto
pre-shaped text keep working untouched.

Mixed content behaves as expected. Latin words and numbers embedded in Arabic
keep their own left to right order, so `"المكتبة Maroto تدعم 42 لغة"` renders with
`Maroto` and `42` readable inside the Arabic.

Line breaking happens on the logical text and each resulting line is shaped and
reordered on its own, with its own base direction, so a wrapped paragraph reads
correctly line by line. Widths are measured on the shaped text, because the
presentation forms have different metrics than the letters they replace and a
lam-alef pair contracts into a single glyph.

## Limitations

* `breakline.DashStrategy` hyphenates, which Arabic does not do, and it breaks
  words at positions measured on the logical characters. Use the default
  `breakline.EmptySpaceStrategy` for Arabic.
* Bidirectional embeddings three levels deep or more, such as an Arabic quotation
  inside a Latin phrase inside an Arabic paragraph, collapse onto the second
  level.
* Only the Arabic letters U+0621 to U+064A are shaped. The Arabic-Extended
  letters used by Persian and Urdu (ڤ, پ, گ) are not yet mapped.

## GoDoc
* [package : rtl](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/rtl)
* [props : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Text)
* [props : Checkbox](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Checkbox)

## Code Example
[filename](../../assets/examples/arabic/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/arabicv2.pdf
```

## Time Execution
[filename](../../assets/text/arabicv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/arabic.json  ':include :type=code')
