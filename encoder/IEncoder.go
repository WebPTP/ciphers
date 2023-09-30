package encoder

/* 二进制数据的编码方式 */
type IEncoder interface {

	/* 算法名称 */
	GetNames() []string

	/* 编码 */
	Encode(data []byte) string

	/* 解码 */
	Decode(data string) ([]byte, error)
}
