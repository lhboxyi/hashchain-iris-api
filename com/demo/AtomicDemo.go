package demo

import (
	"fmt"
	"sync/atomic"
)

/**
保证变量原子性操作
*/

func atomicDemo() {
	i := uint64(10)
	//增加1
	atomic.AddUint64(&i, 1)
	//减少1
	atomic.AddUint64(&i, ^uint64(1))
	fmt.Println(i)
}
