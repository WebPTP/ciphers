package enciphers

import "github.com/WebPTP/ciphers/keypairs"

/* 加密算法 */
type IEncipher interface {

	/* 算法名称 */
	GetNames() []string

	/* 加密数据 */
	Encrypt(key keypairs.Key, plaintext []byte) ([]byte, error)

	/* 解密数据 */
	Decrypt(key keypairs.Key, ciphertext []byte) ([]byte, error)
}
