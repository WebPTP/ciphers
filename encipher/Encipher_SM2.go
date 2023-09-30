package encipher

import (
	"crypto/rand"
	"errors"
	"strconv"

	"github.com/WebPTP/ciphers/keypair"
	"github.com/tjfoc/gmsm/sm2"
)

func NewEncipher_SM2(mode int) *Encipher_SM2 {
	return &Encipher_SM2{
		mode: mode,
	}
}

type Encipher_SM2 struct {
	IEncipher
	mode int
}

func (e Encipher_SM2) GetNames() []string {
	switch e.mode {
	case sm2.C1C3C2:
		return []string{
			"SM2",
			"SM2 C1C3C2",
			"SM2 C1C3C2 DER",
		}
	case sm2.C1C2C3:
		return []string{
			"SM2 C1C2C3",
			"SM2 C1C2C3 DER",
		}
	}
	panic("not supported sm2 mode: " + strconv.Itoa(e.mode))
	// privKey, _ := sm2.GenerateKey(rand.Reader)
}

func (e Encipher_SM2) Encrypt(key keypair.Key, plaintext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PublicKey); ok {
		return sm2.Encrypt(key, plaintext, rand.Reader, e.mode)
	}
	return nil, errors.New("not an *sm2.PublicKey")
}

func (e Encipher_SM2) Decrypt(key keypair.Key, ciphertext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PrivateKey); ok {
		return sm2.Decrypt(key, ciphertext, e.mode)
	}
	return nil, errors.New("not an *sm2.PrivateKey")
}
