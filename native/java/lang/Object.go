package lang

import (
	"jvm-go/native"
	"jvm-go/rtda"
)

func init(){
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class", getClass)
}

func getClass(frame *rtda.Frame){
	// 由于getClas是个实例方法, localVars第一个值是this
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
