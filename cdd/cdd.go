package cdd

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"reflect"
	"time"
)

type Login struct {
	Time time.Time
	Qrcode string
}

type Task struct {
	ctx context.Context
	loginC chan *Login
}

func (t Task) start() {
	// 判断
	ctx, _ := chromedp.NewContext(t.ctx)

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
	if err != nil {
		log.Println(err)
	}

	sel := `#root > div > div > div > main > div > section.login-content.undefined > div > div > div > section > div > div.scan-login.qr-code-activity > div.qr-code`
	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.WaitEnabled(sel),
		chromedp.Screenshot(sel, &buf, chromedp.NodeVisible),
	)
	if err == nil {
		imageBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
		go func() {
			t.loginC<- &Login{
				Time: time.Now(),
				Qrcode: imageBase64,
			}
		}()
	}
	err = chromedp.Run(ctx,
		chromedp.WaitEnabled(`#mms-header-next > div.mms-header-container > div > div.mms-header__list > a:nth-child(1) > div`),
	)
	log.Println(1)
}

func browerEvent(ev interface{}) {
	fmt.Printf("Receive browerEvent TypeName：%v\n", reflect.TypeOf(ev).String())
}

func targetEvent(ev interface{}) {
	fmt.Printf("Receive targetEvent TypeName：%v\n", reflect.TypeOf(ev).String())
}

func (t *Task) Run() {
	t.start()
}

func New(loginC chan *Login) *Task {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.Flag(`disable-extensions`, false),
		chromedp.Flag(`enable-automation`, false),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.109 Safari/537.36"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	jobCtx, _ := chromedp.NewExecAllocator(ctx, options...)

	return &Task{
		ctx: jobCtx,
		loginC: loginC,
	}
}
