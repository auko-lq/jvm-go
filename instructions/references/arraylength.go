package references

import "jvm-go/instructions/base"
import "jvm-go/rtda"

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

// 从操作数栈弹出数组引用, 然后把数组长度压入操作数栈
func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
