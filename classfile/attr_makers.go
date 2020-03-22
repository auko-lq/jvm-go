package classfile

type MarkerAttribute struct{}

// 不建议使用 标记
type DeprecatedAttribute struct{ MarkerAttribute }

// 类成员由编译器生成 标记
type SyntheticAttribute struct{ MarkerAttribute }

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 由于两个只是标记, 没有数据, 所以不用读取
}
