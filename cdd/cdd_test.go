package cdd

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	for i := -4; i > -7; i-- {
		n := time.Now().AddDate(0,0, i)
		d := n.Format("2006-01-02")
		log.Println(d)
	}
}
