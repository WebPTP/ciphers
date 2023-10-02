package encoders

import "encoding/hex"

func NewEncoder_hex() *Encoder_hex {
	return &Encoder_hex{}
}

type Encoder_hex struct {
}

func (*Encoder_hex) GetNames() []string {
	return []string{"HEX"}
}

func (*Encoder_hex) Encode(data []byte) string {
	return hex.EncodeToString(data)
}

func (*Encoder_hex) Decode(data string) ([]byte, error) {
	return hex.DecodeString(data)
}
