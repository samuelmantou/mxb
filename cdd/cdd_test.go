package cdd

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	loginC := make(chan *Login)
	go func() {
		for {
			d := <-loginC
			log.Println(d.Qrcode)
		}
	}()
	task := New(loginC)
	task.Run()

	<-make(chan struct{})
}
