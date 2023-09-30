package encipher

import "github.com/WebPTP/ciphers.go/keypair"

/* 加密算法 */
type IEncipher interface {

	/* 算法名称 */
	GetNames() []string

	/* 加密数据 */
	Encrypt(key keypair.Key, plaintext []byte) ([]byte, error)

	/* 解密数据 */
	Decrypt(key keypair.Key, ciphertext []byte) ([]byte, error)
}
