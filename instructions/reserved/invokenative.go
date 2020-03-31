package reserved

import (
	"github.com/aukocharlie/jvm-go/instructions/base"
	"github.com/aukocharlie/jvm-go/native"
	_ "github.com/aukocharlie/jvm-go/native/java/io"
	_ "github.com/aukocharlie/jvm-go/native/java/lang"
	_ "github.com/aukocharlie/jvm-go/native/sun/misc"
	"github.com/aukocharlie/jvm-go/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
