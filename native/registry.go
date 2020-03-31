package native

import "github.com/aukocharlie/jvm-go/rtda"

// 把本地方法定义为一个函数
// frame就是本地方法的工作空间
// 它是连接JVM和Java类库的桥梁
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	// 类名, 方法名和方法描述符加在一起才能确定唯一的方法
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	// 由于现在自己注册所有本地方法实现
	// 而原本Object等类是通过registerNatives()的本地方法来注册其他本地方法的
	// 现在它没什么用了
	if methodDescriptor == "()V" {
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {}
