package signer

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/WebPTP/ciphers/keypair"
)

type Signer_ECDSA_der struct {
	curve string
}

func NewSigner_ECDSA_der(curve string) *Signer_ECDSA_der {
	return &Signer_ECDSA_der{
		curve: curve,
	}
}

const ECDSA_Curve_P224 = "P224"
const ECDSA_Curve_P256 = "P256"
const ECDSA_Curve_P384 = "P384"
const ECDSA_Curve_P521 = "P521"

func (s Signer_ECDSA_der) GetNames() []string {
	return []string{
		"ECDSA " + s.curve,
		"ECDSA " + s.curve + " DER",
	}
}

func (s Signer_ECDSA_der) Sign(privateKey keypair.PrivateKey, hash []byte) ([]byte, error) {
	return ecdsa.SignASN1(rand.Reader, privateKey.(*ecdsa.PrivateKey), hash)
}

func (s Signer_ECDSA_der) Verify(publicKey keypair.PublicKey, hash []byte, signature []byte) (bool, error) {
	return ecdsa.VerifyASN1(publicKey.(*ecdsa.PublicKey), hash, signature), nil
}
