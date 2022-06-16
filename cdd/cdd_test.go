package cdd

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	for i := 7; i >= 0; i-- {
		e := time.Now().UnixMilli()
		s := e - 86400000*int64(i)
		log.Println(s)
	}
}
