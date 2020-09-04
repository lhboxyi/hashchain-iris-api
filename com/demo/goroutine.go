package demo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
累加1000次
sync.atomic 与 sync.Mutex 对比

sync.atomic 的实现原理大致是向 CPU 发送对某一个块内存的 LOCK 信号，然后就将此内存块加锁，
从而保证了内存块操作的原子性。这种对 CPU 发送信号对内存加锁的方式，比 sync.Mutex 这种在语言层面对内存加锁的方式更底层，因此也更高效。
*/
func MutexGoRoutine() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	count := int64(0)

	//获取当前时间
	startTime := time.Now()

	for i := 0; i < 100000; i++ {
		wg.Add(1)

		go func(i int) {
			mutex.Lock()
			count++
			wg.Done()
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("mutex time:", time.Now().Sub(startTime))
}

func AtomicGoRoutine() {
	var wg sync.WaitGroup
	count := int64(0)

	//获取当前时间
	startTime := time.Now()

	for i := 0; i < 100000; i++ {
		wg.Add(1)

		go func(i int) {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("atomic time:", time.Now().Sub(startTime))
}
