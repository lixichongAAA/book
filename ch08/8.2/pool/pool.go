package main

import (
	"fmt"
	"sync"
	"time"
)

var byteSlicePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	t1 := time.Now().UnixNano()
	//不使用Pool
	for i := 0; i < 100000000; i++ {
		bytes := make([]byte, 1024)
		_ = bytes
	}
	t2 := time.Now().UnixNano()
	//使用Pool
	for i := 0; i < 100000000; i++ {
		bytes := byteSlicePool.Get().(*[]byte)
		_ = bytes
		byteSlicePool.Put(bytes)
	}
	t3 := time.Now().UnixNano()
	fmt.Printf("不使用Pool:%d ns\n", t2-t1)
	fmt.Printf("使用Pool:%d ns\n", t3-t2)
}
