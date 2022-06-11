package main

import (
	cdd2 "mxb/cdd"
	"mxb/jddb"
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
	jddb.Run()
	for {
		time.Sleep(time.Hour * 1)
		task.Reload()
		jddb.Run()
	}
	<-make(chan struct{})
}
