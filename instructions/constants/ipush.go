package constants

import "github.com/aukocharlie/jvm-go/instructions/base"
import "github.com/aukocharlie/jvm-go/rtda"

// 获取byte或short型, 扩展成int型, 再push到操作数栈
type BIPUSH struct {
	val int8
}
type SIPUSH struct {
	val int16
}


func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
