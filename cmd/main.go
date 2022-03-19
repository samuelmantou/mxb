package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"mxb/router"
	"mxb/task"
	"net/http"
	"time"
)

func main() {
	var eg errgroup.Group
	terminalChan := make(chan struct{})
	runChan := make(chan struct{})
	eg.Go(func() error {
		srv := http.Server{
			WriteTimeout: 60 * time.Second,
			ReadTimeout: 60 * time.Second,
			Addr: ":9999",
			Handler: router.Router(runChan, terminalChan),
		}
		if err := srv.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		task.Run(runChan, terminalChan)
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Println(err)
	}
	<-make(chan struct{})
}
