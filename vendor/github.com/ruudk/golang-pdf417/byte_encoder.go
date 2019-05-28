package pdf417

// Code word used to switch to Text mode.
const BYTE_SWITCH_CODE_WORD = 901

// Alternate code word used to switch to Text mode; used when number of
// Texts to encode is divisible by 6.
const BYTE_SWITCH_CODE_WORD_ALT = 924

type ByteEncoder struct {

}

func CreateByteEncoder() *ByteEncoder {
	return new(ByteEncoder)
}

func (ByteEncoder) CanEncode(char string) bool {
	return true
}

func (ByteEncoder) GetSwitchCode(data string) int {
	if (len(data) % 6 == 0) {
		return BYTE_SWITCH_CODE_WORD_ALT
	}

	return BYTE_SWITCH_CODE_WORD
}
