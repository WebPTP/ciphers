package keypair

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"strconv"
)

const RSA_bits_1024 = 1024
const RSA_bits_2048 = 2048
const RSA_bits_3072 = 3072
const RSA_bits_4096 = 4096

type KeypairGenerator_RSA struct {
	IKeypairGenerator
	bits int
}

func NewKeypairGenerator_RSA(bits int) *KeypairGenerator_RSA {
	return &KeypairGenerator_RSA{
		bits: bits,
	}
}

func (k KeypairGenerator_RSA) GetNames() []string {
	return []string{
		"RSA " + strconv.Itoa(k.bits),
		"RSA " + strconv.Itoa(k.bits) + " DER",
	}
}

func (k KeypairGenerator_RSA) GenerateKeyPair() (*IKeypair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, k.bits)
	if err != nil {
		return nil, err
	}
	return &IKeypair{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PublicKey,
	}, nil
}

func (KeypairGenerator_RSA) ExportPublicKey(publicKey PublicKey) ([]byte, error) {
	if publicKey, ok := publicKey.(*rsa.PublicKey); ok {
		return x509.MarshalPKCS1PublicKey(publicKey), nil
	}
	return nil, errors.New("not an *rsa.PublicKey")
}

func (KeypairGenerator_RSA) ImportPublicKey(publicKeyBytes []byte) (PublicKey, error) {
	return x509.ParsePKCS1PublicKey(publicKeyBytes)
}

func (KeypairGenerator_RSA) ExportPrivateKey(privateKey PrivateKey) ([]byte, error) {
	if privateKey, ok := privateKey.(*rsa.PrivateKey); ok {
		return x509.MarshalPKCS1PrivateKey(privateKey), nil
	}
	return nil, errors.New("not an *rsa.PrivateKey")
}

func (KeypairGenerator_RSA) ImportPrivateKey(privateKeyBytes []byte) (PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(privateKeyBytes)
}
