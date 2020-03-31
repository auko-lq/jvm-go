package rtda

import "github.com/aukocharlie/jvm-go/rtda/heap"

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	// program counter
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	// todo: 增加修改栈空间的命令 -Xss
	return &Thread{stack: newStack(1024)}
}

func (self *Thread) PC() int      { return self.pc }
func (self *Thread) SetPC(pc int) { self.pc = pc }

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) PeekFrame() *Frame {
	return self.stack.peek()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}
