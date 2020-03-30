package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

// 用链表实现
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	// 为了弹出时指定top 而记录下一层栈帧
	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		// todo: 异常处理
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

func (self *Stack) peek() *Frame {
	if self._top == nil {
		// todo: 异常处理
		panic("jvm stack is empty!")
	}

	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}
