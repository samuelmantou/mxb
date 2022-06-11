package cdd

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"mxb/task"
	"time"
)

type job struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (j *job) Login(ctx context.Context, data chan<- string) error {
	sel := `#root > div > div > div > main > div > section.login-content.undefined > div > div > div > section > div > div.scan-login.qr-code-activity > div.qr-code`
	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.WaitEnabled(sel),
		chromedp.Sleep(time.Second*3),
		chromedp.Screenshot(sel, &buf, chromedp.NodeVisible),
		chromedp.WaitVisible(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1) > span`),
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				cookies, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					return err
				}

				// 2. 序列化
				cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
				if err != nil {
					return err
				}

				// 3. 存储到临时文件
				if err = ioutil.WriteFile("cookies.tmp", cookiesData, 0755); err != nil {
					return err
				}

				return nil
			}),
		},
	)
	if err == nil {
		imageBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
		res := fmt.Sprintf(`{"task":"cdd", "job":"login", "act":"qrcode", "data":"%s"}`, imageBase64)
		data <- res
	}

	return err
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
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
	)
	return ctx, err
}

func New() task.Job {
	j := job{}
	return &j
}
