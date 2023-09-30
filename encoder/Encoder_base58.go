package encoder

import (
	"errors"

	"github.com/btcsuite/btcutil/base58"
)

func NewEncoder_base58() *Encoder_base58 {
	return &Encoder_base58{}
}

type Encoder_base58 struct {
}

func (Encoder_base58) GetNames() []string {
	return []string{"BASE58"}
}

func (Encoder_base58) Encode(data []byte) string {
	return base58.Encode(data)
}

func (Encoder_base58) Decode(data string) ([]byte, error) {
	if data == "" {
		return make([]byte, 0), nil
	}
	val := base58.Decode(data)
	if len(val) == 0 {
		return nil, errors.New("base58 decoding failed")
	}
	return val, nil
}
