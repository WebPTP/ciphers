package hashers

import "crypto/sha256"

func NewHasher_sha256() *Hasher_sha256 {
	return &Hasher_sha256{}
}

type Hasher_sha256 struct {
	IHasher
}

func (*Hasher_sha256) GetNames() []string {
	return []string{"SHA256"}
}

func (*Hasher_sha256) Digest(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
