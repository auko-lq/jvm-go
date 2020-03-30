package heap

import "jvm-go/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 解析非接口方法符号引用
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		// 接口方法, 即报错
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil{
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

// 寻找method
func lookupMethod(class *Class, name, descriptor string) *Method{
	method := LookupMethodInClass(class, name, descriptor)
	if(method == nil){
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

