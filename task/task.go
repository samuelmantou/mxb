package task

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

type Job interface {
	Navigate(ctx context.Context) (context.Context, error)
	Login()
	Close()
}

type Pipe struct {
	jobs []Job
}

func NewPipe(jobs ...Job) *Pipe {
	p := Pipe{
		jobs: jobs,
	}

	return &p
}

func (receiver *Pipe) Start(ctx context.Context) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.Flag(`disable-extensions`, false),
		chromedp.Flag(`enable-automation`, false),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.109 Safari/537.36"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	jobCtx, cancel := chromedp.NewExecAllocator(ctx, options...)
	defer cancel()

	var err error
	for _, j := range receiver.jobs {
		jobCtx, err = j.Navigate(jobCtx)
		if err != nil {
			log.Println(err)
		}
	}
	for _, j := range receiver.jobs {
		go j.Login()
	}

	<-ctx.Done()
	log.Println("tasks is closing")
}