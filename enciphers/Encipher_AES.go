package enciphers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/WebPTP/ciphers/keypairs"
)

func NewEncipher_AES() *Encipher_AES {
	return &Encipher_AES{}
}

type Encipher_AES struct {
}

func (e Encipher_AES) GetNames() []string {
	return []string{
		"AES",
		"AES CBC",
	}
}

func (Encipher_AES) Encrypt(key keypairs.Key, plaintext []byte) ([]byte, error) {
	if key, ok := key.([]byte); ok {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		iv := make([]byte, aes.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}

		mode := cipher.NewCBCEncrypter(block, iv)

		padding := aes.BlockSize - len(plaintext)%aes.BlockSize
		paddingBytes := bytes.Repeat([]byte{byte(padding)}, padding)
		plaintext = append(plaintext, paddingBytes...)

		ciphertext := make([]byte, len(plaintext))

		mode.CryptBlocks(ciphertext, plaintext)

		ciphertext = append(iv, ciphertext...)
		return ciphertext, nil
	}
	return nil, errors.New("key not an []byte")
}

func (Encipher_AES) Decrypt(key keypairs.Key, ciphertext []byte) ([]byte, error) {
	if key, ok := key.([]byte); ok {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		mode := cipher.NewCBCDecrypter(block, iv)

		plaintext := make([]byte, len(ciphertext))

		mode.CryptBlocks(plaintext, ciphertext)

		padding := int(plaintext[len(plaintext)-1])
		plaintext = plaintext[:len(plaintext)-padding]

		return plaintext, nil
	}
	return nil, errors.New("key not an []byte")
}
