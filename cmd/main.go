package main

import (
	"context"
	"log"
	"mxb/router"
	"mxb/task"
	"mxb/task/cdd"
	consoleJob "mxb/task/console"
	"mxb/task/jd"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	p := task.NewPipe(cdd.New(), jd.New(), consoleJob.New())
	go func(ctx context.Context) {
		p.Start(ctx)
	}(ctx)

	console := router.NewConsole()
	srv := http.Server{
		WriteTimeout: 60 * time.Second,
		ReadTimeout: 60 * time.Second,
		Addr: ":9999",
		Handler: console.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	select {
	case <-sig:
	case <-console.Done():
	}
	srv.Shutdown(context.Background())
	cancel()
	time.Sleep(time.Second * 3)
}
