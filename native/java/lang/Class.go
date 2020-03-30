package lang

import (
	"jvm-go/native"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

const jlClass = "java/lang/Class"

func init(){
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtda.Frame){
	// 先拿到类名字符串引用
	// 由于getPrimitiveClass的参数是个String, 也就意味着它是局部变量表的第一个值
	// 所以下面这句才这么写
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}

// private native String getName0();
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}