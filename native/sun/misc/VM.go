package misc

import (
	"github.com/aukocharlie/jvm-go/instructions/base"
	"github.com/aukocharlie/jvm-go/native"
	"github.com/aukocharlie/jvm-go/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}


// VM类中有个静态块, 会执行initialize
// 这里我们间接地执行initializeSystemClass来初始化
// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
