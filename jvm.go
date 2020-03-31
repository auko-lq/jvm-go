package main

import "fmt"
import "strings"
import "github.com/aukocharlie/jvm-go/classpath"
import "github.com/aukocharlie/jvm-go/instructions/base"
import "github.com/aukocharlie/jvm-go/rtda"
import "github.com/aukocharlie/jvm-go/rtda/heap"

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

// 初始化JVM
func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

// 执行主线程
func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
		return
	}

	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

// 创建参数数组, 即String[] argv
func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JavaString(self.classLoader, arg)
	}
	return argsArr
}
