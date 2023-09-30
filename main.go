package ciphersgo

import (
	"github.com/WebPTP/ciphers.go/encipher"
	"github.com/WebPTP/ciphers.go/encoder"
	"github.com/WebPTP/ciphers.go/hasher"
	"github.com/WebPTP/ciphers.go/keypair"
	"github.com/WebPTP/ciphers.go/signer"
	"github.com/tjfoc/gmsm/sm2"
)

/* 编码器 */
var Encoders map[string]encoder.IEncoder

/* 散列器 */
var Hashers map[string]hasher.IHasher

/* 签名器 */
var Signers map[string]signer.ISigner

/* 秘钥对生成器 */
var KeypairGenerators map[string]keypair.IKeypairGenerator

/* 加密器 */
var Enciphers map[string]encipher.IEncipher

/* 初始化密码相关工具 */
func InitCipher() {
	Encoders = make(map[string]encoder.IEncoder)
	Hashers = make(map[string]hasher.IHasher)
	Signers = make(map[string]signer.ISigner)
	KeypairGenerators = make(map[string]keypair.IKeypairGenerator)
	Enciphers = make(map[string]encipher.IEncipher)

	addEncoder(encoder.NewEncoder_hex())
	addEncoder(encoder.NewEncoder_base58())
	addEncoder(encoder.NewEncoder_base64())

	addHasher(hasher.NewHasher_md5())
	addHasher(hasher.NewHasher_sha256())

	addSigner(signer.NewSigner_ECDSA_der(signer.ECDSA_Curve_P224))
	addSigner(signer.NewSigner_ECDSA_der(signer.ECDSA_Curve_P256))
	addSigner(signer.NewSigner_ECDSA_der(signer.ECDSA_Curve_P384))
	addSigner(signer.NewSigner_ECDSA_der(signer.ECDSA_Curve_P521))

	addKeypairGenerator(keypair.NewKeypairGenerator_ECDSA(keypair.ECDSA_Curve_P224))
	addKeypairGenerator(keypair.NewKeypairGenerator_ECDSA(keypair.ECDSA_Curve_P256))
	addKeypairGenerator(keypair.NewKeypairGenerator_ECDSA(keypair.ECDSA_Curve_P384))
	addKeypairGenerator(keypair.NewKeypairGenerator_ECDSA(keypair.ECDSA_Curve_P521))

	addKeypairGenerator(keypair.NewKeypairGenerator_RSA(keypair.RSA_bits_1024))
	addKeypairGenerator(keypair.NewKeypairGenerator_RSA(keypair.RSA_bits_2048))
	addKeypairGenerator(keypair.NewKeypairGenerator_RSA(keypair.RSA_bits_3072))
	addKeypairGenerator(keypair.NewKeypairGenerator_RSA(keypair.RSA_bits_4096))

	addEncipher(encipher.NewEncipher_RSA(encipher.RSA_bits_1024))
	addEncipher(encipher.NewEncipher_RSA(encipher.RSA_bits_2048))
	addEncipher(encipher.NewEncipher_RSA(encipher.RSA_bits_3072))
	addEncipher(encipher.NewEncipher_RSA(encipher.RSA_bits_4096))

	addEncipher(encipher.NewEncipher_SM2(sm2.C1C2C3))
	addEncipher(encipher.NewEncipher_SM2(sm2.C1C3C2))
	addEncipher(encipher.NewEncipher_SM2_DER())
}

func addEncoder(e encoder.IEncoder) {
	for _, name := range e.GetNames() {
		Encoders[name] = e
	}
}

func addHasher(h hasher.IHasher) {
	for _, name := range h.GetNames() {
		Hashers[name] = h
	}
}

func addSigner(s signer.ISigner) {
	for _, name := range s.GetNames() {
		Signers[name] = s
	}
}

func addKeypairGenerator(k keypair.IKeypairGenerator) {
	for _, name := range k.GetNames() {
		KeypairGenerators[name] = k
	}
}

func addEncipher(e encipher.IEncipher) {
	for _, name := range e.GetNames() {
		Enciphers[name] = e
	}
}
