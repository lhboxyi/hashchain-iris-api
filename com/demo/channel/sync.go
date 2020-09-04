package channel

import (
	"fmt"
	"sync"
)

func sum(count *int, wg *sync.WaitGroup) {
	lock := &sync.Mutex{}
	defer wg.Done()
	for i := 0; i < 101; i++ {
		//如果不使用锁，则每次count的值不一致
		lock.Lock()
		*count += i
		lock.Unlock()
	}
}

/**
 * 启用四个协程取统计count值，如果不加锁，结果会不准
 */
func SyncDemo()  {
	wg := sync.WaitGroup{}
	wg.Add(4)
	count := 0
	go sum(&count, &wg)
	go sum(&count, &wg)
	go sum(&count, &wg)
	go sum(&count, &wg)

	wg.Wait()

	fmt.Println(count)
}
