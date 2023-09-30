package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
)

const ECDSA_Curve_P224 = "P224"
const ECDSA_Curve_P256 = "P256"
const ECDSA_Curve_P384 = "P384"
const ECDSA_Curve_P521 = "P521"

type KeypairGenerator_ECDSA struct {
	IKeypairGenerator
	curve string
}

func NewKeypairGenerator_ECDSA(curve string) *KeypairGenerator_ECDSA {
	return &KeypairGenerator_ECDSA{
		curve: curve,
	}
}

func (k KeypairGenerator_ECDSA) GetNames() []string {
	return []string{
		"ECDSA " + k.curve,
		"ECDSA " + k.curve + " DER",
	}
}

func (KeypairGenerator_ECDSA) GenerateKeyPair() (*IKeypair, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.PublicKey
	return &IKeypair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func (KeypairGenerator_ECDSA) ImportPublicKey(publicKeyBytes []byte) (PublicKey, error) {
	key, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	if key, ok := key.(*ecdsa.PublicKey); ok {
		return key, nil
	}
	return nil, errors.New("not an *ecdsa.PublicKey")
}

func (KeypairGenerator_ECDSA) ExportPublicKey(publicKey PublicKey) ([]byte, error) {
	if publicKey, ok := publicKey.(*ecdsa.PublicKey); ok {
		return x509.MarshalPKIXPublicKey(publicKey)
	}
	return nil, errors.New("not an *ecdsa.PublicKey")
}

func (KeypairGenerator_ECDSA) ImportPrivateKey(privateKeyBytes []byte) (PrivateKey, error) {
	return x509.ParseECPrivateKey(privateKeyBytes)
}

func (KeypairGenerator_ECDSA) ExportPrivateKey(privateKey PrivateKey) ([]byte, error) {
	if privateKey, ok := privateKey.(*ecdsa.PrivateKey); ok {
		return x509.MarshalECPrivateKey(privateKey)
	}
	return nil, errors.New("not an *ecdsa.PrivateKey")
}
