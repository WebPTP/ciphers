package keypair

/* 秘钥 */
type Key any

/* 公钥 */
type PublicKey Key

/* 私钥 */
type PrivateKey Key

/* 非对称加密的秘钥对 */
type IKeypair struct {

	/* 公钥 */
	PublicKey PublicKey

	/* 私钥 */
	PrivateKey PrivateKey
}
