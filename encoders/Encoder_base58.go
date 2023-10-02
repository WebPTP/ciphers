package encoders

import (
	"github.com/mr-tron/base58"
)

func NewEncoder_base58(defaultAlphabet bool, alphabetName string, alphabet *base58.Alphabet) *Encoder_base58 {
	return &Encoder_base58{
		defaultAlphabet: defaultAlphabet,
		alphabetName:    alphabetName,
		alphabet:        alphabet,
	}
}

type Encoder_base58 struct {
	defaultAlphabet bool
	alphabetName    string
	alphabet        *base58.Alphabet
}

func (e Encoder_base58) GetNames() []string {
	name := "BASE58 " + e.alphabetName
	if e.defaultAlphabet {
		return []string{"BASE58", name}
	} else {
		return []string{name}
	}
}

func (e Encoder_base58) Encode(data []byte) string {
	return base58.EncodeAlphabet(data, e.alphabet)
}

func (e Encoder_base58) Decode(data string) ([]byte, error) {
	return base58.DecodeAlphabet(data, e.alphabet)
}
