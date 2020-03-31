package io

import (
	"github.com/aukocharlie/jvm-go/native"
	"github.com/aukocharlie/jvm-go/rtda"
	"os"
	"unsafe"
)

const fos = "java/io/FileOutputStream"

func init() {
	native.Register(fos, "writeBytes", "([BIIZ)V", writeBytes)
}

// todo: 这里的write仅仅是输出到控制台, 而实际的FileOutputStream应该可以输出到文件
// private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
// ([BIIZ)V
func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	//append := vars.GetBoolean(4)

	jBytes := b.Data().([]int8)
	// Java中的byte是有符号类型, 而go中byte则为无符号类型
	// 所以这里需要转换成Go的[]byte变量
	// 再写到控制台
	goBytes := castInt8sToUint8s(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}

// 有符号转为无符号
func castInt8sToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
