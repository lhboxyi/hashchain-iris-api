package channel

import (
	"fmt"
)

func ChannelFlagDone() {
	c := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		fmt.Println("go协程完成")
		c <- true
	}()

	<-c
}
