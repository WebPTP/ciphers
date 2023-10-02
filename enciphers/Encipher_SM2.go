package enciphers

import (
	"crypto/rand"
	"errors"
	"strconv"

	"github.com/WebPTP/ciphers/keypairs"
	"github.com/tjfoc/gmsm/sm2"
)

func NewEncipher_SM2(mode int) *Encipher_SM2 {
	return &Encipher_SM2{
		mode: mode,
	}
}

type Encipher_SM2 struct {
	mode int
}

func (e Encipher_SM2) GetNames() []string {
	switch e.mode {
	case sm2.C1C3C2:
		return []string{
			"SM2",
			"SM2 C1C3C2",
			"SM2 C1C3C2 PKCS8",
		}
	case sm2.C1C2C3:
		return []string{
			"SM2 C1C2C3",
			"SM2 C1C2C3 PKCS8",
		}
	}
	panic("not supported sm2 mode: " + strconv.Itoa(e.mode))
	// privKey, _ := sm2.GenerateKey(rand.Reader)
}

func (e Encipher_SM2) Encrypt(key keypairs.Key, plaintext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PublicKey); ok {
		return sm2.Encrypt(key, plaintext, rand.Reader, e.mode)
	}
	return nil, errors.New("key not an *sm2.PublicKey")
}

func (e Encipher_SM2) Decrypt(key keypairs.Key, ciphertext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PrivateKey); ok {
		return sm2.Decrypt(key, ciphertext, e.mode)
	}
	return nil, errors.New("key not an *sm2.PrivateKey")
}
