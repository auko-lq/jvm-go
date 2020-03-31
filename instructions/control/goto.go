package control

import "github.com/aukocharlie/jvm-go/instructions/base"
import "github.com/aukocharlie/jvm-go/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
