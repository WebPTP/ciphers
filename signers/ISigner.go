package signers

import "github.com/WebPTP/ciphers/keypairs"

type ISigner interface {
	/* 签名器的名称 */
	GetNames() []string

	/* 对数据进行签名 */
	Sign(privateKey keypairs.PrivateKey, hash []byte) ([]byte, error)

	/* 验证签名 */
	Verify(publicKey keypairs.PublicKey, hash []byte, signature []byte) (bool, error)
}
