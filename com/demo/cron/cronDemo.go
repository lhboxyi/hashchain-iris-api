package cron

import "github.com/robfig/cron"

func CronDemo(execDate string) {
	c := cron.New()
	c.AddFunc(execDate, func() {

	})

	c.Start()
	//golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
	select {}
}
