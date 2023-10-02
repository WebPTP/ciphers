package ciphers

import (
	"github.com/WebPTP/ciphers/enciphers"
	"github.com/WebPTP/ciphers/encoders"
	"github.com/WebPTP/ciphers/hashers"
	"github.com/WebPTP/ciphers/keypairs"
	"github.com/WebPTP/ciphers/signers"
	"github.com/mr-tron/base58"
	"github.com/tjfoc/gmsm/sm2"
)

/* 编码器 */
var Encoders map[string]encoders.IEncoder

/* 散列器 */
var Hashers map[string]hashers.IHasher

/* 签名器 */
var Signers map[string]signers.ISigner

/* 秘钥对生成器 */
var KeypairGenerators map[string]keypairs.IKeypairGenerator

/* 加密器 */
var Enciphers map[string]enciphers.IEncipher

/* 初始化密码相关工具 */
func InitCipher() {
	Encoders = make(map[string]encoders.IEncoder)
	Hashers = make(map[string]hashers.IHasher)
	Signers = make(map[string]signers.ISigner)
	KeypairGenerators = make(map[string]keypairs.IKeypairGenerator)
	Enciphers = make(map[string]enciphers.IEncipher)

	addEncoder(encoders.NewEncoder_hex())
	addEncoder(encoders.NewEncoder_base58(true, "BTC", base58.BTCAlphabet))
	addEncoder(encoders.NewEncoder_base58(false, "Flickr", base58.FlickrAlphabet))
	addEncoder(encoders.NewEncoder_base64())

	addHasher(hashers.NewHasher_md5())
	addHasher(hashers.NewHasher_sha256())

	addSigner(signers.NewSigner_ECDSA_der(signers.ECDSA_Curve_P224))
	addSigner(signers.NewSigner_ECDSA_der(signers.ECDSA_Curve_P256))
	addSigner(signers.NewSigner_ECDSA_der(signers.ECDSA_Curve_P384))
	addSigner(signers.NewSigner_ECDSA_der(signers.ECDSA_Curve_P521))

	addKeypairGenerator(keypairs.NewKeypairGenerator_ECDSA(keypairs.ECDSA_Curve_P224))
	addKeypairGenerator(keypairs.NewKeypairGenerator_ECDSA(keypairs.ECDSA_Curve_P256))
	addKeypairGenerator(keypairs.NewKeypairGenerator_ECDSA(keypairs.ECDSA_Curve_P384))
	addKeypairGenerator(keypairs.NewKeypairGenerator_ECDSA(keypairs.ECDSA_Curve_P521))

	addKeypairGenerator(keypairs.NewKeypairGenerator_RSA(keypairs.RSA_bits_1024))
	addKeypairGenerator(keypairs.NewKeypairGenerator_RSA(keypairs.RSA_bits_2048))
	addKeypairGenerator(keypairs.NewKeypairGenerator_RSA(keypairs.RSA_bits_3072))
	addKeypairGenerator(keypairs.NewKeypairGenerator_RSA(keypairs.RSA_bits_4096))

	addKeypairGenerator(keypairs.NewKeypairGenerator_SM2())

	addEncipher(enciphers.NewEncipher_RSA(enciphers.RSA_bits_1024))
	addEncipher(enciphers.NewEncipher_RSA(enciphers.RSA_bits_2048))
	addEncipher(enciphers.NewEncipher_RSA(enciphers.RSA_bits_3072))
	addEncipher(enciphers.NewEncipher_RSA(enciphers.RSA_bits_4096))

	addEncipher(enciphers.NewEncipher_SM2(sm2.C1C2C3))
	addEncipher(enciphers.NewEncipher_SM2(sm2.C1C3C2))
	addEncipher(enciphers.NewEncipher_SM2_DER())
	addEncipher(enciphers.NewEncipher_AES())
}

func addEncoder(e encoders.IEncoder) {
	for _, name := range e.GetNames() {
		Encoders[name] = e
	}
}

func addHasher(h hashers.IHasher) {
	for _, name := range h.GetNames() {
		Hashers[name] = h
	}
}

func addSigner(s signers.ISigner) {
	for _, name := range s.GetNames() {
		Signers[name] = s
	}
}

func addKeypairGenerator(k keypairs.IKeypairGenerator) {
	for _, name := range k.GetNames() {
		KeypairGenerators[name] = k
	}
}

func addEncipher(e enciphers.IEncipher) {
	for _, name := range e.GetNames() {
		Enciphers[name] = e
	}
}
