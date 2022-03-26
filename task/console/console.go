package console

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"mxb/task"
)

type job struct {
	ctx context.Context
	cancel context.CancelFunc
}

func (j *job) Login(ctx context.Context, data chan<- string) error {
	return nil
}

func (j *job) Close() {
	j.cancel()
}

func (j *job) Navigate(ctx context.Context) (context.Context, error) {
	ctx, j.cancel = chromedp.NewContext(ctx)
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate("http://localhost:9999"),
	)
	return ctx, err
}

func New() task.Job {
	j := job{}
	return &j
}