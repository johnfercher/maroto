package pdf417

import (
	"regexp"
	"math"
	"math/big"
)

const NUMBER_SWITCH_CODE_WORD int = 902

type NumberEncoder struct {

}

func CreateNumberEncoder() *NumberEncoder {
	return new(NumberEncoder)
}

func (encoder NumberEncoder) GetName() string {
	return "number"
}

func (encoder NumberEncoder) CanEncode(char string) bool {
	match, err := regexp.MatchString("^[0-9]{1}$", char)

	if err != nil {
		return false
	}

	return match
}

func (encoder NumberEncoder) GetSwitchCode(data string) int {
	return NUMBER_SWITCH_CODE_WORD
}

func (encoder NumberEncoder) Encode(digits string, addSwitchCode bool) []int {
	digitCount := len(digits)
	chunkCount := int(math.Ceil(float64(digitCount) / float64(44)))

	codeWords := []int{}

	if (addSwitchCode) {
		codeWords = append(codeWords, NUMBER_SWITCH_CODE_WORD)
	}

	for i := 0; i < chunkCount; i++ {
		start := i * 44
		end := start + 44
		if end > digitCount {
			end = digitCount
		}
		chunk := digits[start:end]

		cws := encodeChunk(chunk)

		codeWords = append(codeWords, cws...)
	}

	return codeWords
}

func encodeChunk(chunkInput string) []int {
	chunk := big.NewInt(0)

	_, ok := chunk.SetString("1" + chunkInput, 10)

	if ! ok {
		panic("Failed converting")
	}

	cws := []int{}

	for chunk.Cmp(big.NewInt(0)) > 0 {
		newChunk, cw := chunk.DivMod(
			chunk,
			big.NewInt(900),
			big.NewInt(0),
		)

		chunk = newChunk

		cws = append([]int{int(cw.Int64())}, cws...)
	}

	return cws
}
