// Package rtl renders right to left text, Arabic in particular, in the form a
// PDF writer can draw.
//
// PDF text operators draw glyphs one after another from left to right and do
// not implement any script specific logic. Arabic therefore needs two
// transformations before it reaches the writer:
//
//   - shaping, which replaces each letter by the contextual presentation form
//     that connects to its neighbours, and contracts the mandatory lam-alef
//     ligatures;
//   - reordering, which lays the logical order text out in the visual order the
//     writer must emit.
//
// Without them Arabic comes out as disconnected letters in reversed order.
package rtl

// Process returns text ready to be drawn by a left to right writer.
//
// Text without Arabic characters is returned unchanged, byte for byte, so that
// calling Process on arbitrary content is safe.
//
// Process expects a single already wrapped line. Running it on a paragraph
// before the line breaking would reorder the text as one unit and produce
// visually reversed lines, so callers must break lines first and process each
// resulting line on its own.
func Process(text string) string {
	if !ContainsArabic(text) {
		return text
	}

	return reorder(shape(text))
}

// ContainsArabic reports whether text holds at least one character from the
// Arabic blocks, including the presentation forms.
func ContainsArabic(text string) bool {
	for _, r := range text {
		switch {
		case r >= 0x0600 && r <= 0x06FF, // Arabic
			r >= 0x0750 && r <= 0x077F, // Arabic Supplement
			r >= 0x08A0 && r <= 0x08FF, // Arabic Extended-A
			r >= 0xFB50 && r <= 0xFDFF, // Arabic Presentation Forms-A
			r >= 0xFE70 && r <= 0xFEFF: // Arabic Presentation Forms-B
			return true
		}
	}

	return false
}
