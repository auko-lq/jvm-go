package extended

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

// Branch if reference is null
// 判断引用是否为null
type IFNULL struct{ base.BranchInstruction }

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }


func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
