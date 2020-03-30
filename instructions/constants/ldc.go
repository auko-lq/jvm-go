package constants

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

// Push item from run-time constant pool
// 从运行时常量池中加载常量值, 并推入操作数栈
type LDC struct{ base.Index8Instruction }

// Push item from run-time constant pool (wide index)
// 与LDC的区别在于操作数的宽度
type LDC_W struct{ base.Index16Instruction }

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JavaString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		// 加载类对象字面值
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	// todo: MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
