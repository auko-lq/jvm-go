package comparisons

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

// Branch if reference comparison succeeds
// 判断引用是否相等
type IF_ACMPEQ struct{ base.BranchInstruction }

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
