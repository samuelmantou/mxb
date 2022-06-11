package main

import (
	"log"
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
	for {
		time.Sleep(time.Hour * 1)
		task.Reload()
		go func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
			jddb.Run()
		}()
	}
	<-make(chan struct{})
}
