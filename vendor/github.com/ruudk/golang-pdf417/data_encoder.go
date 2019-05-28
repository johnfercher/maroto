package pdf417

type Encoder interface {
	GetName() string
	CanEncode(char string) bool
	GetSwitchCode(data string) int
	Encode(data string, addSwitchCode bool) []int
}

type DataEncoder struct {
	Encoders []Encoder
	DefaultEncoder Encoder
}

type Chain struct {
	Data string
	Encoder Encoder
}

func CreateDataEncoder() DataEncoder {
	numberEncoder := CreateNumberEncoder()
	textEncoder := CreateTextEncoder()

	encoder := DataEncoder{
		[]Encoder{numberEncoder, textEncoder},
		textEncoder,
	}

	return encoder
}

func (dataEncoder DataEncoder) Encode(data string) []int {
	chains := dataEncoder.SplitToChains(data)

	if len(chains) == 0 {
		panic("hmmm")
	}


	firstEncoder := chains[0].Encoder
	addSwitchCode := firstEncoder.GetName() != dataEncoder.DefaultEncoder.GetName()

	cws := []int{}

	for _, chainWithEncoder := range chains {
		encoded := chainWithEncoder.Encoder.Encode(
			chainWithEncoder.Data,
			addSwitchCode,
		)

		cws = append(cws, encoded...)

		addSwitchCode = true
	}

	return cws
}

func (dataEncoder DataEncoder) SplitToChains(data string) []Chain {
	chains := []Chain{}
	chainData := ""
	encoder := dataEncoder.DefaultEncoder

	for i := 0; i < len(data); i++ {
		char := string(data[i])

		newEncoder := getEncoder(
			dataEncoder.Encoders,
			char,
		)

		if newEncoder != encoder {
			if len(chainData) > 0 {
				chains = append(chains, Chain{chainData, encoder})
				chainData = ""
			}

			encoder = newEncoder
		}

		chainData = chainData + char
	}

	if len(chainData) > 0 {
		chains = append(chains, Chain{chainData, encoder})
	}

	return chains
}

func getEncoder(encoders []Encoder, char string) Encoder {
	for _, encoder := range encoders {
		if encoder.CanEncode(char) {
			return encoder
		}
	}

	panic("Cannot encode character " + char)
}
