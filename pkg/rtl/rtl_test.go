package rtl_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/rtl"
	"github.com/stretchr/testify/assert"
)

// The expected values are Arabic Presentation Forms-B characters (U+FE70 to
// U+FEFF), not the letters of the input: that block is what a font draws to
// connect the letters to each other. They read as garbled Arabic in an editor,
// which is expected, because they are stored in the visual order the PDF
// writer emits them and each letter carries its contextual form. The comments
// name the form picked for each letter so the tables can be reviewed against
// the Unicode chart.

func TestProcess_Shaping(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			// meem initial, hah medial, meem medial, dal final.
			name:     "when letters connect on both sides, should use the medial forms",
			input:    "محمد",
			expected: "ﺪﻤﺤﻣ",
		},
		{
			// kaf initial, teh medial, alef final, beh isolated. The alef is
			// right joining, so the beh that follows it cannot connect back.
			name:     "when a right joining letter is in the middle, should isolate the next letter",
			input:    "كتاب",
			expected: "ﺏﺎﺘﻛ",
		},
		{
			// sheen initial, yeh final, hamza isolated.
			name:     "when the word ends in hamza, should keep the hamza non joining",
			input:    "شيء",
			expected: "ﺀﻲﺷ",
		},
		{
			// alef isolated, lam initial, lam medial, heh final.
			name:     "when an alef precedes a lam, should not form a ligature",
			input:    "الله",
			expected: "ﻪﻠﻟﺍ",
		},
		{
			// ain initial, lam medial, alef maksura final.
			name:     "when the word ends in alef maksura, should use the final form",
			input:    "على",
			expected: "ﻰﻠﻋ",
		},
		{
			// meem initial, dal final, reh isolated, seen initial, teh marbuta final.
			name:     "when the word ends in teh marbuta, should use the final form",
			input:    "مدرسة",
			expected: "ﺔﺳﺭﺪﻣ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, rtl.Process(tt.input))
		})
	}
}

func TestProcess_LamAlefLigatures(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "when lam is followed by alef and nothing precedes, should use the isolated ligature",
			input:    "لا",
			expected: "ﻻ",
		},
		{
			name:     "when lam is followed by alef with madda, should use the isolated ligature",
			input:    "لآ",
			expected: "ﻵ",
		},
		{
			name:     "when lam is followed by alef with hamza above, should use the isolated ligature",
			input:    "لأ",
			expected: "ﻷ",
		},
		{
			name:     "when lam is followed by alef with hamza below, should use the isolated ligature",
			input:    "لإ",
			expected: "ﻹ",
		},
		{
			// beh initial, then the ligature in its final form.
			name:     "when a dual joining letter precedes the ligature, should use the final ligature",
			input:    "بلا",
			expected: "ﻼﺑ",
		},
		{
			name:     "when a dual joining letter precedes the madda ligature, should use the final ligature",
			input:    "بلآ",
			expected: "ﻶﺑ",
		},
		{
			name:     "when a dual joining letter precedes the hamza above ligature, should use the final ligature",
			input:    "بلأ",
			expected: "ﻸﺑ",
		},
		{
			name:     "when a dual joining letter precedes the hamza below ligature, should use the final ligature",
			input:    "بلإ",
			expected: "ﻺﺑ",
		},
		{
			// The fatha sits between the lam and the alef and must not stop the
			// contraction; it stays attached after the ligature.
			name:     "when a mark sits between lam and alef, should still ligate and keep the mark",
			input:    "لَا",
			expected: "ﻻَ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, rtl.Process(tt.input))
		})
	}
}

func TestProcess_Diacritics(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			// dal final, meem initial, then the shadda and the fatha in the
			// order they were given.
			name:     "when a letter carries a shadda and a fatha, should keep both in order",
			input:    "مَّد",
			expected: "ﺪﻣَّ",
		},
		{
			// The tanween sits on the last letter of the word, right before the
			// space, and must not be moved across it.
			name:     "when a tanween ends a word before a space, should keep it on that word",
			input:    "كتابً ك",
			expected: "ﻙ ﺏًﺎﺘﻛ",
		},
		{
			// The sukun is transparent: the noon still connects to the dal.
			name:     "when a sukun sits between two letters, should not break the joining",
			input:    "منْد",
			expected: "ﺪﻨْﻣ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, rtl.Process(tt.input))
		})
	}
}

// TestProcess_AllDiacritics sweeps the whole U+064B..U+0652 range. Every mark
// has to be transparent for the joining, so the beh stays initial and the dal
// stays final, and the mark has to stay right after the beh it decorates.
func TestProcess_AllDiacritics(t *testing.T) {
	t.Parallel()

	for mark := rune(0x064B); mark <= 0x0652; mark++ {
		t.Run(string(mark), func(t *testing.T) {
			t.Parallel()
			input := "ب" + string(mark) + "د"

			assert.Equal(t, "ﺪﺑ"+string(mark), rtl.Process(input))
		})
	}
}

// TestProcess_IsolatedRepertoire walks the 36 letters U+0621..U+063A and
// U+0641..U+064A on their own, which must give the isolated presentation form
// of each one.
func TestProcess_IsolatedRepertoire(t *testing.T) {
	t.Parallel()

	isolated := map[string]string{
		"ء": "ﺀ", "آ": "ﺁ", "أ": "ﺃ", "ؤ": "ﺅ",
		"إ": "ﺇ", "ئ": "ﺉ", "ا": "ﺍ", "ب": "ﺏ",
		"ة": "ﺓ", "ت": "ﺕ", "ث": "ﺙ", "ج": "ﺝ",
		"ح": "ﺡ", "خ": "ﺥ", "د": "ﺩ", "ذ": "ﺫ",
		"ر": "ﺭ", "ز": "ﺯ", "س": "ﺱ", "ش": "ﺵ",
		"ص": "ﺹ", "ض": "ﺽ", "ط": "ﻁ", "ظ": "ﻅ",
		"ع": "ﻉ", "غ": "ﻍ", "ف": "ﻑ", "ق": "ﻕ",
		"ك": "ﻙ", "ل": "ﻝ", "م": "ﻡ", "ن": "ﻥ",
		"ه": "ﻩ", "و": "ﻭ", "ى": "ﻯ", "ي": "ﻱ",
	}

	assert.Len(t, isolated, 36)

	for letter, expected := range isolated {
		t.Run(letter, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, expected, rtl.Process(letter))
		})
	}
}

func TestProcess_MixedRuns(t *testing.T) {
	t.Parallel()

	const marhaba = "ﺎﺒﺣﺮﻣ"

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			// The base direction is right to left, so the Latin run is laid out
			// to the left of the Arabic one.
			name:     "when arabic leads and latin follows, should put the latin run first",
			input:    "مرحبا Hello",
			expected: "Hello " + marhaba,
		},
		{
			// The base direction is left to right here, so the run order is kept.
			name:     "when latin leads, should keep the run order and reverse only the arabic",
			input:    "Hi مرحبا",
			expected: "Hi " + marhaba,
		},
		{
			// Digits are not strong, so the base direction still comes from the
			// Arabic and the number ends up on the right.
			name:     "when digits lead, should take the base direction from the arabic",
			input:    "42 مرحبا",
			expected: marhaba + " 42",
		},
		{
			name:     "when digits are embedded, should keep them in left to right order",
			input:    "مرحبا 42 عام",
			expected: "ﻡﺎﻋ 42 " + marhaba,
		},
		{
			// Arabic-Indic digits read left to right just like the ASCII ones.
			name:     "when arabic indic digits are embedded, should keep them in left to right order",
			input:    "عمر ٤٢",
			expected: "٤٢ ﺮﻤﻋ",
		},
		{
			name:     "when a decimal is embedded, should not split it around the separator",
			input:    "قيمة 3.14 نهاية",
			expected: "ﺔﻳﺎﻬﻧ 3.14 ﺔﻤﻴﻗ",
		},
		{
			// The two Latin words keep their order relative to each other.
			name:     "when a multi word latin run is embedded, should keep its internal order",
			input:    "في New York الان",
			expected: "ﻥﻻﺍ New York ﻲﻓ",
		},
		{
			name:     "when letters and digits are embedded together, should keep the token intact",
			input:    "مر ABC123 مر",
			expected: "ﺮﻣ ABC123 ﺮﻣ",
		},
		{
			name:     "when the text ends in an arabic question mark, should place it on the left",
			input:    "ما؟",
			expected: "؟ﺎﻣ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, rtl.Process(tt.input))
		})
	}
}

func TestProcess_NonArabicPassthrough(t *testing.T) {
	t.Parallel()

	inputs := []string{
		"",
		"Hello, World!",
		"123 456",
		"3.14",
		"Ünïcödé àçcents",
		"line with  double  spaces",
		"symbols !@#$%^&*()",
		"日本語",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, input, rtl.Process(input), "non arabic input must be returned byte for byte")
		})
	}
}

func TestContainsArabic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "when the text is empty, should be false", input: "", expected: false},
		{name: "when the text is latin, should be false", input: "Hello", expected: false},
		{name: "when the text is digits, should be false", input: "42", expected: false},
		{name: "when the text is arabic, should be true", input: "مرحبا", expected: true},
		{name: "when the text mixes arabic and latin, should be true", input: "Hi مرحبا", expected: true},
		{name: "when the text holds arabic indic digits, should be true", input: "٤٢", expected: true},
		{name: "when the text is already shaped, should be true", input: "ﻻ", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, rtl.ContainsArabic(tt.input))
		})
	}
}
