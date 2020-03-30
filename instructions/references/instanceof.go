package references

import "jvm-go/instructions/base"
import "jvm-go/rtda"
import "jvm-go/rtda/heap"

// Determine if object is of given type
// 判断对象是否是某个类的实例
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 先弹出对象引用
	ref := stack.PopRef()
	if ref == nil {
		// 如果对象时null, 则将0压入操作数栈, 表示结果为false
		stack.PushInt(0)
		return
	}

	// 去获取类符号引用, 进行判断
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
