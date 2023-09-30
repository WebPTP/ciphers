package hasher

import "crypto/md5"

func NewHasher_md5() *Hasher_md5 {
	return &Hasher_md5{}
}

type Hasher_md5 struct {
}

func (*Hasher_md5) GetNames() []string {
	return []string{"MD5"}
}

func (*Hasher_md5) Digest(data []byte) []byte {
	hasher := md5.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
