package task

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"mxb/task/cdd"
	"mxb/task/jd"
)

func Run(runChan <-chan struct{}, terminalChan <-chan struct{}) {
	var err error
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.Flag(`disable-extensions`, false),
		chromedp.Flag(`enable-automation`, false),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.109 Safari/537.36"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	cddCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	err = chromedp.Run(cddCtx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate("https://mms.pinduoduo.com"),
	)

	if err != nil {
		fmt.Println(err)
	}

	jdCtx, cancel := chromedp.NewContext(cddCtx)
	defer cancel()

	err = chromedp.Run(jdCtx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate("https://procurement.jd.com/procurement/initListPage"),
	)

	ctrCtx, cancel := chromedp.NewContext(jdCtx)
	defer cancel()
	err = chromedp.Run(ctrCtx,
		chromedp.Navigate("http://localhost:9999"),
	)

	if err != nil {
		fmt.Println(err)
	}

	go func(cddCtx context.Context, jdCtx context.Context) {
		for {
			select {
			case _, ok := <-runChan:
				if !ok {
					return
				}
				err = chromedp.Run(cddCtx,
					cdd.Task(),
				)
				err = chromedp.Run(jdCtx,
					jd.Task(),
				)
			}
		}
	}(cddCtx, jdCtx)
	<-terminalChan
}
