package classpath

import(
	"os"
	"path/filepath"
)

// 类加载器加载目标
type Classpath struct{
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 依次从多种路径中搜索class文件
func(self *Classpath) ReadClass(className string)([]byte, Entry, error){
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil{
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil{
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// 返回用户类路径
func(self *Classpath) String()string {
	return self.userClasspath.String()
}

// 解析启动类和扩展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib","ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户类路径
func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption == ""{
		// 没指定的话就是当前目录
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// 获取jre路径, 优先使用-Xjre选项, 没有的话就使用当前目录, 还没有就用JAVA_HOME
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh !=""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool{
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

