package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
// 三种情况下, S类型引用值可以赋给T类型(S转为T)
// - 同类
// - T是S父类
// - S实现T接口

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				// 当S是接口, T不是接口
				// 那么只有T是Object时, S才能转为T
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				// 数组可以强转为Object
				return t.isJlObject()
			} else {
				// t is interface
				// 数组也可以强转为Cloneable和Serializable
				// 因为数组实现了这两个接口
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			// 如果两个都是数组, 就判断他们是否为同一基本类型
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			// 或者如果数组元素是引用类型, 就判断该引用类型是否可以强转
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
	return false
}

// self extends c
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}
