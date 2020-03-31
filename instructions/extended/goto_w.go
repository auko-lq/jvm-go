package extended

import (
	"github.com/aukocharlie/jvm-go/instructions/base"
	"github.com/aukocharlie/jvm-go/rtda"
)

// Branch always (wide index)
// 和goto的区别在于索引从2byte变成了4byte
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
