package enciphers

import (
	"crypto/rand"
	"errors"

	"github.com/WebPTP/ciphers/keypairs"
	"github.com/tjfoc/gmsm/sm2"
)

func NewEncipher_SM2_DER() *Encipher_SM2_DER {
	return &Encipher_SM2_DER{}
}

type Encipher_SM2_DER struct {
}

func (Encipher_SM2_DER) GetNames() []string {
	return []string{
		"SM2 DER",
	}
}

func (Encipher_SM2_DER) Encrypt(key keypairs.Key, plaintext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PublicKey); ok {
		return sm2.EncryptAsn1(key, plaintext, rand.Reader)
	}
	return nil, errors.New("key not an *sm2.PublicKey")
}

func (Encipher_SM2_DER) Decrypt(key keypairs.Key, ciphertext []byte) ([]byte, error) {
	if key, ok := key.(*sm2.PrivateKey); ok {
		return sm2.DecryptAsn1(key, ciphertext)
	}
	return nil, errors.New("key not an *sm2.PrivateKey")
}
