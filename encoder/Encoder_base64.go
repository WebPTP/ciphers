package encoder

import "encoding/base64"

func NewEncoder_base64() *Encoder_base64 {
	return &Encoder_base64{}
}

type Encoder_base64 struct {
}

func (*Encoder_base64) GetNames() []string {
	return []string{"BASE64"}
}

func (*Encoder_base64) Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (*Encoder_base64) Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
