package references

import "jvm-go/instructions/base"
import "jvm-go/rtda"
import "jvm-go/rtda/heap"

// 创建对象
type NEW struct{ base.Index16Instruction }

// 从运行时常量池中拿到类符号引用
// 解析符号引用来获取类数据
// 依此创建对象
func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		// 规范规定, 接口和抽象类若要实例化则要抛出异常
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
