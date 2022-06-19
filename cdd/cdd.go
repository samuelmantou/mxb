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
	Time   time.Time
	Qrcode string
}

type Task struct {
	ctx    context.Context
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

type OrderResultListResp struct {
	ProductId          int64  `json:"productId"`
	WarehouseProductId int64  `json:"warehouseProductId"`
	ProductName        string `json:"productName"`
	ProductThumbUrl    string `json:"productThumbUrl"`
	Total              int    `json:"total"`
	SellUnitTotal      int    `json:"sellUnitTotal"`
	SellUnitName       string `json:"sellUnitName"`
	QuantityManageInfo struct {
		HitConfirmInboundGray      bool `json:"hitConfirmInboundGray"`
		HitInboundDisplayGray      bool `json:"hitInboundDisplayGray"`
		ShowSellPlan               bool `json:"showSellPlan"`
		ReplaceInboundStat         bool `json:"replaceInboundStat"`
		ShowRealTimeWarehouseStock bool `json:"showRealTimeWarehouseStock"`
		QuantityInfo               struct {
			Quantity               int `json:"quantity"`
			SellUnitQuantity       int `json:"sellUnitQuantity"`
			InboundTotal           int `json:"inboundTotal"`
			PurchasingInboundTotal int `json:"purchasingInboundTotal"`
		} `json:"quantityInfo"`
		ConfirmEventInfo          interface{} `json:"confirmEventInfo"`
		IsWarehouseStockSync      bool        `json:"isWarehouseStockSync"`
		ShowInventoryDetail       bool        `json:"showInventoryDetail"`
		EnableInventoryDistribute bool        `json:"enableInventoryDistribute"`
		InventoryDetail           struct {
			WarehouseDistributableInventory interface{} `json:"warehouseDistributableInventory"`
			WarehouseTotalQuantity          interface{} `json:"warehouseTotalQuantity"`
			CenterSalableInventory          int         `json:"centerSalableInventory"`
			CenterInventoryList             []struct {
				WarehouseId                   int         `json:"warehouseId"`
				WarehouseName                 string      `json:"warehouseName"`
				CenterSupplierInboundQuantity int         `json:"centerSupplierInboundQuantity"`
				CenterInventory               interface{} `json:"centerInventory"`
				CenterAllotInboundQuantity    int         `json:"centerAllotInboundQuantity"`
				CenterAllotOnWayQuantity      int         `json:"centerAllotOnWayQuantity"`
				DisplayAllotTotal             int         `json:"displayAllotTotal"`
				DisplayRealTimeInventory      interface{} `json:"displayRealTimeInventory"`
			} `json:"centerInventoryList"`
			ShareDistributableInventory   interface{} `json:"shareDistributableInventory"`
			ShareRelatedInventory         interface{} `json:"shareRelatedInventory"`
			ShareDistributionList         interface{} `json:"shareDistributionList"`
			ProcessDistributableInventory interface{} `json:"processDistributableInventory"`
			ProcessRelatedInventory       interface{} `json:"processRelatedInventory"`
			ProcessDistributionList       interface{} `json:"processDistributionList"`
		} `json:"inventoryDetail"`
		AppointmentData        interface{} `json:"appointmentData"`
		ShowAdvanceDeliverData bool        `json:"showAdvanceDeliverData"`
		AdvanceDeliverData     interface{} `json:"advanceDeliverData"`
	} `json:"quantityManageInfo"`
	SpecQuantityDetails []struct {
		SpecDesc         string `json:"specDesc"`
		Total            int    `json:"total"`
		SellUnitCount    int    `json:"sellUnitCount"`
		SellUnitTotal    int    `json:"sellUnitTotal"`
		Quantity         int    `json:"quantity"`
		SellUnitQuantity int    `json:"sellUnitQuantity"`
		PriceDetail      []struct {
			SupplierPrice       interface{} `json:"supplierPrice"`
			SupplierPriceDetail interface{} `json:"supplierPriceDetail"`
			Total               int         `json:"total"`
			SellUnitTotal       int         `json:"sellUnitTotal"`
			SessionEnd          bool        `json:"sessionEnd"`
			LastUpdateTime      interface{} `json:"lastUpdateTime"`
			NextUpdateTime      interface{} `json:"nextUpdateTime"`
			StatId              interface{} `json:"statId"`
		} `json:"priceDetail"`
	} `json:"specQuantityDetails"`
	HasCode69        bool   `json:"hasCode69"`
	HasPrintedLabel  bool   `json:"hasPrintedLabel"`
	SessionDate      int64  `json:"sessionDate"`
	SessionEnd       bool   `json:"sessionEnd"`
	AreaId           int    `json:"areaId"`
	AreaName         string `json:"areaName"`
	WarehouseId      int    `json:"warehouseId"`
	WarehouseName    string `json:"warehouseName"`
	WarehouseAddress string `json:"warehouseAddress"`
	WarehouseGroupId int    `json:"warehouseGroupId"`
	IsProduct        bool   `json:"isProduct"`
	HasMultiSchedule bool   `json:"hasMultiSchedule"`
	SalesPlan        struct {
		PlanSales           int         `json:"planSales"`
		LatestPlanSendTime  int64       `json:"latestPlanSendTime"`
		StockSynced         int         `json:"stockSynced"`
		SalesPlanDetailList interface{} `json:"salesPlanDetailList"`
	} `json:"salesPlan"`
	ProductSaleStatus interface{} `json:"productSaleStatus"`
}

type OrderResp struct {
	Success   bool        `json:"success"`
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  interface{} `json:"errorMsg"`
	Result    struct {
		Total      int `json:"total"`
		UpdateTime struct {
			Now            int64       `json:"now"`
			SessionEnd     bool        `json:"sessionEnd"`
			LastUpdateTime interface{} `json:"lastUpdateTime"`
			NextUpdateTime interface{} `json:"nextUpdateTime"`
			AutoUpdateTime interface{} `json:"autoUpdateTime"`
		} `json:"updateTime"`
		ResultList             []OrderResultListResp `json:"resultList"`
		ErrorMsgList           interface{}           `json:"errorMsgList"`
		ShowSellPlan           bool                  `json:"showSellPlan"`
		ShowInventoryDetail    bool                  `json:"showInventoryDetail"`
		ShowAdvanceDeliverData bool                  `json:"showAdvanceDeliverData"`
	} `json:"result"`
}

type ShouHouReq struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	AreaId    int    `json:"areaId"`
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
}

type ShouHouResp struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type OrderBaseReq struct {
	headers network.Headers
	orderParams OrderReq
	shouHouParams ShouHouReq
}

var (
	orderReqs = make(map[int]OrderBaseReq)
	shouhouReqs = make(map[int]OrderBaseReq)
	orderUrl = "https://mc.pinduoduo.com/cartman-mms/orderManagement/pageQueryDetail"
	shouhouUrl = "https://mc.pinduoduo.com/ragnaros-mms/after/sales/manage/queryProductAfterSalesStatistic"
)

func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*network.EventRequestWillBeSent); ok {
			req := event.Request
			if req.URL == orderUrl {
				var o OrderReq
				if err := json.Unmarshal([]byte(req.PostData), &o); err != nil {
					log.Println(err)
				}
				orderReqs[o.AreaId] = OrderBaseReq{
					headers: req.Headers,
					orderParams: o,
				}
			}
			if req.URL == shouhouUrl {
				go func(req network.Request) {
					var o ShouHouReq
					if err := json.Unmarshal([]byte(req.PostData), &o); err != nil {
						log.Println(err)
					}
					for _, areaId := range []int{4, 19881230, 19881231} {
						shouhouReqs[areaId] = OrderBaseReq{
							headers: req.Headers,
							shouHouParams: o,
						}
					}
					log.Println("售后结束抓包")
				}(*req)
			}
		}
	})
}

var listeners bool

func (t *Task) Reload() {
	var err error
	ctx := t.ctx
	if listeners == false {
		listenForNetworkEvent(ctx)
		listeners = true
	}
	err = chromedp.Run(ctx,
		t.log("订单页面初始化开始"),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second*3),
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				cs, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					return err
				}
				for _, c := range cs {
					cookies = append(cookies, &http.Cookie{
						Name:   c.Name,
						Value:  c.Value,
						Path:   c.Path,
						Domain: c.Domain,
					})
				}

				return nil
			}),
		},
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second*3),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建爬虫开始"),
		chromedp.Sleep(time.Second*3),
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
		chromedp.Sleep(time.Second*3),
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

	err = chromedp.Run(ctx,
		t.log("粤东爬虫开始"),
		chromedp.Sleep(time.Second*3),
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
		chromedp.Sleep(time.Second*3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/after-sales-manage"),
		chromedp.Sleep(time.Second*3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/after-sales-manage"),
		chromedp.Sleep(time.Second*3),
		t.log("售后页面初始化结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > section > button.BTN_outerWrapper_1nl4rhj.BTN_primary_1nl4rhj.BTN_medium_1nl4rhj.BTN_outerWrapperBtn_1nl4rhj`),
		t.log("福建省爬虫查询"),
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("该轮爬虫结束")
	log.Println("开始爬取数据")
	t.grapData()
	log.Println("结束爬取数据")

	t.ctx = ctx
}

func (t *Task) grapData() {
	for k, item := range orderReqs {
		log.Printf("订单地区:%d开始抓包", k)
		time.Sleep(time.Second * 5)
		for d := 7; d >= 0; d-- {
			s := time.Now().UnixMilli() - 86400000*int64(d)
			e := s
			for i := 1; i < 3; i++ {
				o := item.orderParams
				o.Page = i
				o.PageSize = 100
				o.StartSessionTime = s
				o.EndSessionTime = e
				b, _ := json.Marshal(o)
				postReq, err := http.NewRequest(http.MethodPost, orderUrl, bytes.NewBuffer(b))
				if err != nil {
					log.Println(err)
				}
				c := &http.Client{}
				for k, r := range item.headers {
					postReq.Header.Set(k, fmt.Sprintf("%s", r))
				}
				for _, cookie := range cookies {
					postReq.AddCookie(cookie)
				}
				resp, err := c.Do(postReq)
				if err != nil {
					log.Println(err)
					log.Println("发送请求销售地址失败")
					return
				}
				defer resp.Body.Close()

				body, _ := ioutil.ReadAll(resp.Body)
				var r OrderResp
				json.Unmarshal(body, &r)
				if r.Result.Total == 0 || len(r.Result.ResultList) == 0 {
					continue
				}
				date := time.UnixMilli(r.Result.ResultList[0].SessionDate).Format("2006-01-02")
				http.PostForm("http://api.mxb.j1mi.com/index/transfer?use=ajax", url.Values{
					"data": {string(body)},
					"from": {"order"},
					"date": {date},
					"params": {string(b)},
				})
				time.Sleep(time.Second * 2)
			}
		}
		log.Printf("订单地区: %d结束抓包", k)
	}
	for k, item := range shouhouReqs {
		log.Printf("售后地区:%d开始抓包", k)
		for i := 1; i < 6; i++ {
			o := item.shouHouParams
			o.AreaId = k
			o.PageNum = i
			o.PageSize = 30
			for j := -1; j > -7; j-- {
				n := time.Now().AddDate(0, 0, j)
				d := n.Format("2006-01-02")
				o.StartDate = d
				o.EndDate = d
				b, _ := json.Marshal(o)

				postReq, err := http.NewRequest(http.MethodPost, shouhouUrl, bytes.NewBuffer(b))
				if err != nil {
					log.Println(err)
				}
				c := &http.Client{}
				for k, r := range item.headers {
					postReq.Header.Set(k, fmt.Sprintf("%s", r))
				}
				for _, c := range cookies {
					postReq.AddCookie(c)
				}
				resp, err := c.Do(postReq)
				if err != nil {
					log.Println(err)
					log.Println("发送请求售后地址失败")
					return
				}
				defer resp.Body.Close()

				body, _ := ioutil.ReadAll(resp.Body)
				var r ShouHouResp
				json.Unmarshal(body, &r)
				if r.Success == false {
					continue
				}
				http.PostForm("http://api.mxb.j1mi.com/index/transfer?use=ajax", url.Values{
					"data": {string(body)},
					"from": {"shouhou"},
					"date": {d},
					"params": {string(b)},
				})
			}
		}
		log.Printf("售后地区:%d结束抓包", k)
	}
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
		chromedp.Sleep(time.Second*2),
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
		ctx:    jobCtx,
		loginC: loginC,
	}
}
