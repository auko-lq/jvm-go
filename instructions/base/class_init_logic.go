package base

import (
	"github.com/aukocharlie/jvm-go/rtda"
	"github.com/aukocharlie/jvm-go/rtda/heap"
)

// jvms 5.5
// 类初始化就是执行类的初始化方法<clinit>
// - 执行new指令创建类实例, 但类还没初始化
// - 执行putstatic, getstatic指令存取类的静态变量, 但类还没初始化
// - 执行invokestatic调用静态方法, 但类还没初始化
// - 子类要初始化, 父类要先初始化
// - 反射操作
func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	// 递归地把clinit方法压到栈顶, 保证父类的初始化方法在上面
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
