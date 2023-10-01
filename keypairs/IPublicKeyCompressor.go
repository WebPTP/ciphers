package keypairs

/* 公钥压缩器 */
type IPublicKeyCompressor interface {

	/* 压缩公钥 */
	CompressPublicKey(publicKey PublicKey) ([]byte, error)

	/* 解压公钥 */
	DecompressPublicKey(data []byte) (PublicKey, error)
}
