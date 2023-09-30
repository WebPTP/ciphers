package encipher

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"strconv"

	"github.com/WebPTP/ciphers.go/keypair"
)

const RSA_bits_1024 = 1024
const RSA_bits_2048 = 2048
const RSA_bits_3072 = 3072
const RSA_bits_4096 = 4096

func NewEncipher_RSA(bits int) *Encipher_RSA {
	return &Encipher_RSA{
		bits: bits,
	}
}

type Encipher_RSA struct {
	IEncipher
	bits int
}

func (e Encipher_RSA) GetNames() []string {
	return []string{
		"RSA " + strconv.Itoa(e.bits),
		"RSA " + strconv.Itoa(e.bits) + " DER",
	}
}

func (e Encipher_RSA) Encrypt(key keypair.Key, plaintext []byte) ([]byte, error) {
	if publicKey, ok := key.(*rsa.PublicKey); ok {
		if publicKey.Size() != e.bits {
			return nil, errors.New("public key size mismatch")
		}
		return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	}
	return nil, errors.New("not an *rsa.PublicKey")
}

func (e Encipher_RSA) Decrypt(key keypair.Key, ciphertext []byte) ([]byte, error) {
	if privateKey, ok := key.(*rsa.PrivateKey); ok {
		if privateKey.Size() != e.bits {
			return nil, errors.New("private key size mismatch")
		}
		return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	}
	return nil, errors.New("not an *rsa.PrivateKey")
}
