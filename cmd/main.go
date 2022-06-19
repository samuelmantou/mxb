package main

import (
	"log"
	cdd2 "mxb/cdd"
	"mxb/jddb"
	"time"
)

func main() {
	log.Println("v1.0.0 爬虫启动：注意，只有在服务器上才可以爬取金蝶数据。")
	loginC := make(chan *cdd2.Login)
	go func() {
		for {
			<-loginC
		}
	}()
	task := cdd2.New(loginC, false)
	task.Run()
	task.Reload()
	jddb.Run()
	for {
		time.Sleep(time.Minute * 5)
		task.Reload()
		jddb.Run()
	}
	<-make(chan struct{})
}
