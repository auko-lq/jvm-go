package references

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

// Check whether object is of given type
// 这条指令通常用于类型转换
// 与instanceof类似, 区别在于它不会改变操作数栈状态
// 如果判断失败, 不是压入0, 而是抛出ClassCastException异常
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	// 再压回去
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
