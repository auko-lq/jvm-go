package rtda

import "github.com/aukocharlie/jvm-go/rtda/heap"

type Slot struct{
	num int32
	ref *heap.Object
}