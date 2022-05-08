package main

import (
	cdd2 "mxb/cdd"
	"time"
)

func main() {
	loginC := make(chan *cdd2.Login)
	go func() {
		for {
			<-loginC
		}
	}()
	task := cdd2.New(loginC, false)
	task.Run()
	task.Reload()
	for {
		time.Sleep(time.Hour * 5)
		task.Reload()
	}
	<-make(chan struct{})
}
