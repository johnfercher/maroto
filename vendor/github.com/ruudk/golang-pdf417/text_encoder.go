package pdf417

// Code word used to switch to Text mode.
const TEXT_SWITCH_CODE_WORD int = 900

// Since each code word consists of 2 characters, a padding value is
// needed when encoding a single character. 29 is used as padding because
// it's a switch in all 4 submodes, and doesn't add any data.
const PADDING_VALUE = 29

// Uppercase submode.
const SUBMODE_UPPER = "SUBMODE_UPPER"

// Lowercase submode.
const SUBMODE_LOWER = "SUBMODE_LOWER"

// mixed submode (numbers and some punctuation).
const SUBMODE_MIXED = "SUBMODE_MIXED"

// Punctuation submode.
const SUBMODE_PUNCT = "SUBMODE_PUNCT"

// Switch to uppercase submode.
const SWITCH_UPPER = "SWITCH_UPPER"

// Switch to uppercase submode for a single character.
const SWITCH_UPPER_SINGLE = "SWITCH_UPPER_SINGLE"

// Switch to lowercase submode.
const SWITCH_LOWER = "SWITCH_LOWER"

// Switch to mixed submode.
const SWITCH_MIXED = "SWITCH_MIXED"

// Switch to punctuation submode.
const SWITCH_PUNCT = "SWITCH_PUNCT"

// Switch to punctuation submode for a single character.
const SWITCH_PUNCT_SINGLE = "SWITCH_PUNCT_SINGLE"

type TextEncoder struct {
	CharacterTables map[string][]string
	Switching       map[string]map[string][]string
	SwitchSubmode   map[string]string
	ReverseLookup   map[string]map[string]int
}

func CreateTextEncoder() *TextEncoder {
	encoder := new(TextEncoder)
	encoder.CharacterTables = map[string][]string{
		SUBMODE_UPPER: []string{
			"A", "B", "C", "D", "E", "F", "G", "H", "I",
			"J", "K", "L", "M", "N", "O", "P", "Q", "R",
			"S", "T", "U", "V", "W", "X", "Y", "Z", " ",
			SWITCH_LOWER,
			SWITCH_MIXED,
			SWITCH_PUNCT_SINGLE,
		},

		SUBMODE_LOWER: []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i",
			"j", "k", "l", "m", "n", "o", "p", "q", "r",
			"s", "t", "u", "v", "w", "x", "y", "z", " ",
			SWITCH_UPPER_SINGLE,
			SWITCH_MIXED,
			SWITCH_PUNCT_SINGLE,
		},

		SUBMODE_MIXED: []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8",
			"9", "&", "\r", "\t", ",", ":", "#", "-", ".",
			"$", "/", "+", "%", "*", "=", "^",
			SWITCH_PUNCT, " ",
			SWITCH_LOWER,
			SWITCH_UPPER,
			SWITCH_PUNCT_SINGLE,
		},

		SUBMODE_PUNCT: []string{
			";", "<", ">", "@", "[", "\\", "]", "_", "`",
			"~", "!", "\r", "\t", ",", ":", "\n", "-", ".",
			"$", "/", "\"", "|", "*", "(", ")", "?", "{", "}", "'",
			SWITCH_UPPER,
		},
	}

	encoder.Switching = map[string]map[string][]string{
		SUBMODE_UPPER: map[string][]string{
			SUBMODE_LOWER: []string{SWITCH_LOWER},
			SUBMODE_MIXED: []string{SWITCH_MIXED},
			SUBMODE_PUNCT: []string{SWITCH_MIXED, SWITCH_PUNCT},
		},

		SUBMODE_LOWER: map[string][]string{
			SUBMODE_UPPER: []string{SWITCH_MIXED, SWITCH_UPPER},
			SUBMODE_MIXED: []string{SWITCH_MIXED},
			SUBMODE_PUNCT: []string{SWITCH_MIXED, SWITCH_PUNCT},
		},

		SUBMODE_MIXED: map[string][]string{
			SUBMODE_UPPER: []string{SWITCH_UPPER},
			SUBMODE_LOWER: []string{SWITCH_LOWER},
			SUBMODE_PUNCT: []string{SWITCH_PUNCT},
		},

		SUBMODE_PUNCT: map[string][]string{
			SUBMODE_UPPER: []string{SWITCH_UPPER},
			SUBMODE_LOWER: []string{SWITCH_UPPER, SWITCH_LOWER},
			SUBMODE_MIXED: []string{SWITCH_UPPER, SWITCH_MIXED},
		},
	}

	encoder.SwitchSubmode = map[string]string{
		SWITCH_UPPER: SUBMODE_UPPER,
		SWITCH_LOWER: SUBMODE_LOWER,
		SWITCH_PUNCT: SUBMODE_PUNCT,
		SWITCH_MIXED: SUBMODE_MIXED,
	}

	encoder.ReverseLookup = make(map[string]map[string]int)
	for submode, codes := range encoder.CharacterTables {
		for row, char := range codes {
			if encoder.ReverseLookup[char] == nil {
				encoder.ReverseLookup[char] = make(map[string]int)
			}

			encoder.ReverseLookup[char][submode] = int(row)
		}
	}

	return encoder
}

func (encoder TextEncoder) GetName() string {
	return "text"
}

func (encoder TextEncoder) CanEncode(char string) bool {
	return encoder.ReverseLookup[char] != nil
}

func (TextEncoder) GetSwitchCode(data string) int {
	return TEXT_SWITCH_CODE_WORD
}

func (encoder TextEncoder) Encode(data string, addSwitchCode bool) []int {
	interim := encodeinterim(encoder, data);

	return encodeFinal(interim, addSwitchCode)
}

func encodeinterim(encoder TextEncoder, data string) []int {
	submode := SUBMODE_UPPER

	codes := []int{}

	for i := 0; i < len(data); i++ {
		char := string(data[i])

		if (existsInSubmode(encoder, char, submode) == false) {
			prevSubmode := submode

			submode = getSubmode(encoder, char)

			switchCodes := getSwitchCodes(encoder, prevSubmode, submode)

			codes = append(codes, switchCodes...)
		}

		codes = append(
			codes,
			getCharacterCode(encoder, char, submode),
		)
	}

	return codes
}

func getSubmode(encoder TextEncoder, char string) string {
	_, ok := encoder.ReverseLookup[char]

	if ! ok {
		panic("Weird, not found")
	}

	for key := range encoder.ReverseLookup[char] {
		return key
	}

	return "INVALID"
}

func getSwitchCodes(encoder TextEncoder, from string, to string) []int {
	switches := encoder.Switching[from][to]

	codes := []int{}

	for _, switcher := range switches {
		codes = append(codes, getCharacterCode(encoder, switcher, from))

		from = encoder.SwitchSubmode[switcher]
	}

	return codes
}

func encodeFinal(codes []int, addSwitchCode bool) []int {
	codeWords := []int{}

	if addSwitchCode {
		codeWords = append(codeWords, TEXT_SWITCH_CODE_WORD)
	}

	chunks := [][]int{}
	chunkPart := []int{}
	i := 0
	for _, code := range codes {
		chunkPart = append(chunkPart, code)

		i++

		if i % 2 == 0 {
			chunks = append(chunks, chunkPart)

			chunkPart = []int{}
		}
	}

	if len(chunkPart) > 0 {
		chunks = append(chunks, chunkPart)
	}

	for _, chunk := range chunks {
		if len(chunk) == 1 {
			chunk = append(chunk, PADDING_VALUE)
		}

		codeWords = append(
			codeWords,
			30 * chunk[0] + chunk[1],
		)
	}

	return codeWords
}

func getCharacterCode(encoder TextEncoder, char string, submode string) int {
	cw, ok := encoder.ReverseLookup[char][submode]

	if ! ok {
		panic("This is not possible")
	}

	return cw
}

func existsInSubmode(encoder TextEncoder, char string, submode string) bool {
	_, ok := encoder.ReverseLookup[char][submode]

	return ok
}
