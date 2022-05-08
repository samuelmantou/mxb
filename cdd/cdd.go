package cdd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

type OrderReq struct {
	Page             int   `json:"page"`
	PageSize         int   `json:"pageSize"`
	AreaId           int   `json:"areaId"`
	WarehouseIds     []int `json:"warehouseIds"`
	StartSessionTime int64 `json:"startSessionTime"`
	EndSessionTime   int64 `json:"endSessionTime"`
}

type ShouHouReq struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	AreaId    int    `json:"areaId"`
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
}

func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*network.EventRequestWillBeSent); ok {
			req := event.Request
			if req.URL == "https://mc.pinduoduo.com/cartman-mms/orderManagement/pageQueryDetail" {
				go func(req network.Request) {
					time.Sleep(time.Second * 5)
					b := []byte(req.PostData)

					var o OrderReq
					if err := json.Unmarshal(b, &o); err != nil {
						log.Println(err)
					}
					o.PageSize = 100
					e := time.Now().UnixMilli()
					s := e - 86400000 * 7
					o.StartSessionTime = s
					o.EndSessionTime = e
					b, _ = json.Marshal(o)

					postReq, err := http.NewRequest(http.MethodPost, req.URL, bytes.NewBuffer(b))
					if err != nil {
						log.Println(err)
					}
					c := &http.Client{}
					for k, r := range req.Headers {
						postReq.Header.Set(k, fmt.Sprintf("%s", r))
					}
					for _, c := range cookies {
						postReq.AddCookie(c)
					}
					resp, err := c.Do(postReq)
					defer resp.Body.Close()

					body, _ := ioutil.ReadAll(resp.Body)

					http.PostForm("http://api.mxb.j1mi.com/index/transfer?use=ajax", url.Values{
						"data": {string(body)},
						"from": {"order"},
						"date": {""},
					})
				}(*req)
			}
			if req.URL == "https://mc.pinduoduo.com/ragnaros-mms/after/sales/manage/queryProductAfterSalesStatistic" {
				go func(req network.Request) {
					time.Sleep(time.Second * 5)
					b := []byte(req.PostData)

					var o ShouHouReq
					if err := json.Unmarshal(b, &o); err != nil {
						log.Println(err)
					}
					o.PageSize = 30
					for i := -4; i > -7; i-- {
						n := time.Now().AddDate(0,0, i)
						d := n.Format("2006-01-02")

						o.StartDate = d
						o.EndDate = d
						b, _ = json.Marshal(o)

						postReq, err := http.NewRequest(http.MethodPost, req.URL, bytes.NewBuffer(b))
						if err != nil {
							log.Println(err)
						}
						c := &http.Client{}
						for k, r := range req.Headers {
							postReq.Header.Set(k, fmt.Sprintf("%s", r))
						}
						for _, c := range cookies {
							postReq.AddCookie(c)
						}
						resp, err := c.Do(postReq)
						defer resp.Body.Close()

						body, _ := ioutil.ReadAll(resp.Body)

						http.PostForm("http://api.mxb.tech/index/transfer?use=ajax", url.Values{
							"data": {string(body)},
							"from": {"shouhou"},
							"date": {d},
						})
					}
				}(*req)
			}
		}
	})
}

func (t *Task) Reload() {
	var err error
	ctx := t.ctx

	listenForNetworkEvent(ctx)

	err = chromedp.Run(ctx,
		t.log("订单页面初始化开始"),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				cs, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					return err
				}
				for _, c := range cs {
					cookies = append(cookies, &http.Cookie{
						Name: c.Name,
						Value: c.Value,
						Path: c.Path,
						Domain: c.Domain,
					})
				}

				return nil
			}),
		},
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 3),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("福建爬虫结束"),
	)
	if err != nil {
		log.Println("福建爬虫失败:" + err.Error())
	}

	err = chromedp.Run(ctx,
		t.log("粤西爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(2)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("粤西爬虫结束"),
	)
	if err != nil {
		log.Println("粤西爬虫失败:" + err.Error())
	}
	//

	err = chromedp.Run(ctx,
		t.log("粤东爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(3)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("粤东爬虫结束"),
	)
	if err != nil {
		log.Println("粤东爬虫失败:" + err.Error())
	}

	err = chromedp.Run(ctx,
		t.log("售后页面初始化开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/after-sales-manage"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/after-sales-manage"),
		chromedp.Sleep(time.Second * 3),
		t.log("售后页面初始化结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("福建省爬虫查询"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(2)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("福建省爬虫查询"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(3)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("福建省爬虫查询"),
	)
	if err != nil {
		log.Println(err)
	}

	t.ctx = ctx
}

func (t *Task) log(flag string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		log.Println(flag)
		return nil
	}
}

var cookies []*http.Cookie

func (t *Task) start() {
	// 判断
	ctx, _ := chromedp.NewContext(t.ctx)

	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Sleep(time.Second),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		chromedp.WaitEnabled(`#mms-header-next > div.mms-header-container > div > div.mms-header__list > a:nth-child(1) > div`),
		chromedp.Sleep(time.Second * 2),
	)
	if err != nil {
		log.Println("进入首页失败:" + err.Error())
	}

	t.ctx = ctx
}

func (t *Task) Run() {
	t.start()
}

func New(loginC chan *Login, headless bool) *Task {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
		chromedp.WindowSize(1920, 1024),
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
