package rtl

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/bidi"
)

// reorder turns a logical order string into the visual order a left to right
// writer such as the PDF text operator must emit to display it correctly.
//
// Run segmentation and the resolution of neutral characters (spaces,
// punctuation, brackets) and of numbers follow UAX#9 and are delegated to
// golang.org/x/text/unicode/bidi, which is already part of the module
// dependency tree. The reordering itself is applied here because
// bidi.Ordering exposes its runs in logical order and only reports the
// direction of each run, not its embedding level: golang.org/x/text computes
// the levels but discards everything except their parity, and the functions
// that implement the UAX#9 rule L2 reordering are unexported.
//
// The consequence is a two level reordering model: runs that follow the base
// direction and runs that oppose it. That is exact for the mixed content this
// package targets (Arabic with embedded Latin words, numbers and punctuation).
// Embeddings three levels deep or more, such as an Arabic quotation nested
// inside a Latin phrase which is itself nested inside an Arabic paragraph,
// collapse onto the second level and may be placed incorrectly.
func reorder(text string) string {
	var paragraph bidi.Paragraph

	_, err := paragraph.SetString(text)
	if err != nil {
		return text
	}

	ordering, err := paragraph.Order()
	if err != nil {
		return text
	}

	runs := make([]string, 0, ordering.NumRuns())
	for i := range ordering.NumRuns() {
		run := ordering.Run(i)
		if run.Direction() == bidi.RightToLeft {
			runs = append(runs, reverseClusters(run.String()))
			continue
		}
		runs = append(runs, run.String())
	}

	if baseDirection(text) == bidi.RightToLeft {
		reverseSlice(runs)
	}

	var out strings.Builder
	for _, run := range runs {
		out.WriteString(run)
	}

	return out.String()
}

// baseDirection resolves the paragraph direction the same way the UAX#9 rules
// P2 and P3 do: the direction of the first strong character, defaulting to
// left to right when the text holds none. Digits, spaces and punctuation are
// not strong and are therefore skipped.
func baseDirection(text string) bidi.Direction {
	for _, r := range text {
		properties, _ := bidi.LookupRune(r)

		class := properties.Class()
		if class == bidi.L {
			return bidi.LeftToRight
		}

		if class == bidi.R || class == bidi.AL {
			return bidi.RightToLeft
		}
	}

	return bidi.LeftToRight
}

// reverseClusters reverses text keeping every combining mark attached to, and
// positioned after, the base character it decorates. Reversing rune by rune
// would move the marks onto the wrong character.
func reverseClusters(text string) string {
	var (
		out     []rune
		current []rune
	)

	flush := func() {
		out = append(current, out...)
		current = nil
	}

	for _, r := range text {
		if unicode.Is(unicode.Mn, r) && len(current) > 0 {
			current = append(current, r)
			continue
		}
		flush()
		current = []rune{r}
	}
	flush()

	return string(out)
}

func reverseSlice(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
