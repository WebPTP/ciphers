package enciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"strconv"

	"github.com/WebPTP/ciphers/keypairs"
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
	bits int
}

func (e Encipher_RSA) GetNames() []string {
	return []string{
		"RSA " + strconv.Itoa(e.bits),
		"RSA " + strconv.Itoa(e.bits) + " DER",
	}
}

func (e Encipher_RSA) Encrypt(key keypairs.Key, plaintext []byte) ([]byte, error) {
	if publicKey, ok := key.(*rsa.PublicKey); ok {
		if publicKey.N.BitLen() != e.bits {
			return nil, errors.New("public key size mismatch: publicKey[" + strconv.Itoa(publicKey.N.BitLen()) + "] e.bits[" + strconv.Itoa(e.bits) + "]")
		}
		return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	}
	return nil, errors.New("key not an *rsa.PublicKey")
}

func (e Encipher_RSA) Decrypt(key keypairs.Key, ciphertext []byte) ([]byte, error) {
	if privateKey, ok := key.(*rsa.PrivateKey); ok {
		if privateKey.N.BitLen() != e.bits {
			return nil, errors.New("private key size mismatch")
		}
		return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	}
	return nil, errors.New("key not an *rsa.PrivateKey")
}
