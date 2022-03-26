package main

import (
	"context"
	"log"
	"mxb/router"
	"mxb/task"
	"mxb/task/cdd"
	"mxb/ws"
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

	wsPool := ws.New()

	//p := task.NewPipe(cdd.New(), jd.New())
	p := task.NewPipe(wsPool.Send(), cdd.New())
	go func(ctx context.Context) {
		p.Start(ctx)
	}(ctx)

	console := router.NewConsole(wsPool)
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
	time.Sleep(time.Second)
}
