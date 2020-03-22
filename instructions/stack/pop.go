package stack

import "jvm-go/instructions/base"
import "jvm-go/rtda"

// double和long需要pop2弹出两个slot
type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }


func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
