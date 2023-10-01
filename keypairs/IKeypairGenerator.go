package keypairs

/* 秘钥对生成器 */
type IKeypairGenerator interface {

	/* 获取生成器名称 */
	GetNames() []string

	/* 生成秘钥对 */
	GenerateKeyPair() (*IKeypair, error)

	/* 导入公钥 */
	ImportPublicKey(publicKeyBytes []byte) (PublicKey, error)

	/* 导出公钥 */
	ExportPublicKey(publicKey PublicKey) ([]byte, error)

	/* 导入私钥 */
	ImportPrivateKey(privateKeyBytes []byte) (PrivateKey, error)

	/* 导出私钥 */
	ExportPrivateKey(privateKey PrivateKey) ([]byte, error)
}
