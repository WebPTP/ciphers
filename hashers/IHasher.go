package hashers

type IHasher interface {
	/* 算法名称 */
	GetNames() []string

	/* 执行数据摘要 */
	Digest(data []byte) []byte
}
