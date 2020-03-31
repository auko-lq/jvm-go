package lang

import (
	"github.com/aukocharlie/jvm-go/native"
	"github.com/aukocharlie/jvm-go/rtda"
	"unsafe"
)

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass();
func getClass(frame *rtda.Frame) {
	// 由于getClass是个实例方法, localVars第一个值是this
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	// unsafe获取引用的指针, 转成uintptr
	// uintptr是指针的数字表达, 它足够容纳任何指针的位模式
	// 再转成int32就可以得到一个引用的唯一hash
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
	//  获取对象的this引用, 即类
	this := frame.LocalVars().GetThis()

	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	// 检查该类是否可以clone
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}
