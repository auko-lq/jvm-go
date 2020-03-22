package classfile

// 具体指向哪种常量因字段类型而异
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
}

func(self *ConstantValueAttribute) ConstantValueIndex() uint16{
	return self.constantValueIndex
}