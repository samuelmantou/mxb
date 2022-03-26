package cdd

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"mxb/task"
	"time"
)

type job struct {
	ctx context.Context
	cancel context.CancelFunc
}

func (j *job) Login() {
	sel := `#root > div > div > div > main > div > section.login-content.undefined > div > div > div > div.login-tab > div > div.tab-item.last-item`
	chromedp.Run(j.ctx,
		chromedp.WaitEnabled(sel),
		chromedp.Sleep(time.Second * 1),
		chromedp.Click(sel),
		chromedp.SendKeys(`#usernameId`, "18620978045", chromedp.ByID),
	)
}

func (j *job) Close() {
	j.cancel()
}

func (j *job) Navigate(ctx context.Context) (context.Context, error) {
	ctx, j.cancel = chromedp.NewContext(ctx)
	j.ctx = ctx
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate("https://mms.pinduoduo.com"),
	)
	return ctx, err
}

func New() task.Job {
	j := job{}
	return &j
}
