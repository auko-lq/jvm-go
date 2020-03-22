package math

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

// Subtract double
type DSUB struct{ base.NoOperandsInstruction }

// Subtract float
type FSUB struct{ base.NoOperandsInstruction }

// Subtract int
type ISUB struct{ base.NoOperandsInstruction }

// Subtract long
type LSUB struct{ base.NoOperandsInstruction }


func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
