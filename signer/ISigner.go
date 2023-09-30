package signer

import "github.com/WebPTP/ciphers.go/keypair"

type ISigner interface {
	/* 签名器的名称 */
	GetNames() []string

	/* 对数据进行签名 */
	Sign(privateKey keypair.PrivateKey, hash []byte) ([]byte, error)

	/* 验证签名 */
	Verify(publicKey keypair.PublicKey, hash []byte, signature []byte) (bool, error)
}
