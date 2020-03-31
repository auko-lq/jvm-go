package constants

import (
	"github.com/aukocharlie/jvm-go/instructions/base"
	"github.com/aukocharlie/jvm-go/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {}