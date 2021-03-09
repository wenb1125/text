package chain

/**
 * 迭代器的接口标准声明：
 * 		①判断是否还要数据
 * 		②取出下一个数据
 */
type ChainIterator interface {
	HasNext() bool
	Next() Block
}
