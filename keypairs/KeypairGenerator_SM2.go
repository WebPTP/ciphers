package keypairs

import (
	"crypto/rand"
	"errors"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func NewKeypairGenerator_SM2() *KeypairGenerator_SM2 {
	return &KeypairGenerator_SM2{}
}

type KeypairGenerator_SM2 struct {
	IKeypairGenerator
	IPublicKeyCompressor
}

func (KeypairGenerator_SM2) GetNames() []string {
	return []string{
		"SM2",
		"SM2 C1C3C2",
		"SM2 C1C3C2 PKCS8",
		"SM2 C1C2C3",
		"SM2 C1C2C3 PKCS8",
	}
}

func (KeypairGenerator_SM2) GenerateKeyPair() (*IKeypair, error) {
	key, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &IKeypair{
		PrivateKey: key,
		PublicKey:  &key.PublicKey,
	}, nil
}

func (KeypairGenerator_SM2) ImportPublicKey(publicKeyBytes []byte) (PublicKey, error) {
	return x509.ParseSm2PublicKey(publicKeyBytes)
}

func (KeypairGenerator_SM2) ExportPublicKey(publicKey PublicKey) ([]byte, error) {
	if publicKey, ok := publicKey.(*sm2.PublicKey); ok {
		return x509.MarshalSm2PublicKey(publicKey)
	}
	return nil, errors.New("not an *sm2.PublicKey")
}

func (KeypairGenerator_SM2) ImportPrivateKey(privateKeyBytes []byte) (PrivateKey, error) {
	// PKCS8
	return x509.ParsePKCS8UnecryptedPrivateKey(privateKeyBytes)
}

func (KeypairGenerator_SM2) ExportPrivateKey(privateKey PrivateKey) ([]byte, error) {
	if privateKey, ok := privateKey.(*sm2.PrivateKey); ok {
		// PKCS8
		return x509.MarshalSm2UnecryptedPrivateKey(privateKey)
	}
	return nil, errors.New("not an *sm2.PrivateKey")
}

func (KeypairGenerator_SM2) CompressPublicKey(publicKey PublicKey) ([]byte, error) {
	if publicKey, ok := publicKey.(*sm2.PublicKey); ok {
		return sm2.Compress(publicKey), nil
	}
	return nil, errors.New("not an *sm2.PublicKey")
}

func (KeypairGenerator_SM2) DecompressPublicKey(data []byte) (publicKey PublicKey, err error) {
	defer func() {
		if e := recover(); e != nil {
			publicKey = nil
			err = e.(error)
		}
	}()
	publicKey = sm2.Decompress(data)
	err = nil
	return
}
