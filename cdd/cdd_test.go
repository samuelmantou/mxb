package cdd

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	loginC := make(chan *Login)
	go func() {
		for {
			d := <-loginC
			log.Println(d.Qrcode)
		}
	}()
	task := New(loginC, false)
	task.Run()
	time.Sleep(time.Second * 5)
	task.reload()
	<-make(chan struct{})
}
