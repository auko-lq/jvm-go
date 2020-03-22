package comparisons

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

// Compare float
/*
	浮点数的比较不同于int, 因为浮点数有infinity和NaN两个特殊值
	如1.0f/0.0f 和 0.0f/0.0f
	分别对应infinity和NaN
	所以除了大于等于小于外, 还要加上第四种比较
*/
type FCMPG struct{ base.NoOperandsInstruction }

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
