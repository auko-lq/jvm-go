package lang

import "github.com/aukocharlie/jvm-go/native"
import "github.com/aukocharlie/jvm-go/rtda"
import "github.com/aukocharlie/jvm-go/rtda/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// 将字符串放入字符串常量池
// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
