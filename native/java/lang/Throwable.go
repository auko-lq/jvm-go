package lang

import (
	"fmt"
	"github.com/aukocharlie/jvm-go/native"
	"github.com/aukocharlie/jvm-go/rtda"
	"github.com/aukocharlie/jvm-go/rtda/heap"
)

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

const jlThrowable = "java/lang/Throwable"

func init() {
	native.Register(jlThrowable, "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

// 创建一个调用链
func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 因为调用fillInStackTrace时, 栈顶已经有些多余的栈帧了, 需要跳过
	// 后面的 + 2 表示跳过栈顶的fillInStackTrace()和fillInStackTrace(int)
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 由于此时栈中正在执行异常类的构造函数, 所以需要计算他们有多少层, 然后跳过
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}


func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}