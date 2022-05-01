package cdd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"time"
)

type OrderResponse struct {
	Success   bool        `json:"success"`
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  interface{} `json:"errorMsg"`
	Result    struct {
		Total      int `json:"total"`
		UpdateTime struct {
			Now            int64 `json:"now"`
			SessionEnd     bool  `json:"sessionEnd"`
			LastUpdateTime int64 `json:"lastUpdateTime"`
			NextUpdateTime int64 `json:"nextUpdateTime"`
			AutoUpdateTime int64 `json:"autoUpdateTime"`
		} `json:"updateTime"`
		ResultList []struct {
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
					WarehouseDistributableInventory int `json:"warehouseDistributableInventory"`
					WarehouseTotalQuantity          int `json:"warehouseTotalQuantity"`
					CenterSalableInventory          int `json:"centerSalableInventory"`
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
					ShareDistributableInventory int `json:"shareDistributableInventory"`
					ShareRelatedInventory       int `json:"shareRelatedInventory"`
					ShareDistributionList       []struct {
						WarehouseId            int    `json:"warehouseId"`
						WarehouseName          string `json:"warehouseName"`
						DistributableInventory int    `json:"distributableInventory"`
						Inventory              int    `json:"inventory"`
						IsRelatedWarehouse     bool   `json:"isRelatedWarehouse"`
					} `json:"shareDistributionList"`
					ProcessDistributableInventory int `json:"processDistributableInventory"`
					ProcessRelatedInventory       int `json:"processRelatedInventory"`
					ProcessDistributionList       []struct {
						WarehouseId            int    `json:"warehouseId"`
						WarehouseName          string `json:"warehouseName"`
						DistributableInventory int    `json:"distributableInventory"`
						Inventory              int    `json:"inventory"`
						IsRelatedWarehouse     bool   `json:"isRelatedWarehouse"`
					} `json:"processDistributionList"`
				} `json:"inventoryDetail"`
				AppointmentData struct {
					AppointmentTotal        int `json:"appointmentTotal"`
					CenterAppointmentDetail struct {
						WarehouseId   int    `json:"warehouseId"`
						WarehouseName string `json:"warehouseName"`
						Quantity      int    `json:"quantity"`
					} `json:"centerAppointmentDetail"`
					ShareAppointmentDetails interface{} `json:"shareAppointmentDetails"`
				} `json:"appointmentData"`
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
					SupplierPrice       int `json:"supplierPrice"`
					SupplierPriceDetail struct {
						RealSupplierPrice   int         `json:"realSupplierPrice"`
						ShareWarehousePrice int         `json:"shareWarehousePrice"`
						DamagePrice         int         `json:"damagePrice"`
						TransportPrice      interface{} `json:"transportPrice"`
						SupplierPrice       int         `json:"supplierPrice"`
					} `json:"supplierPriceDetail"`
					Total          int         `json:"total"`
					SellUnitTotal  int         `json:"sellUnitTotal"`
					SessionEnd     bool        `json:"sessionEnd"`
					LastUpdateTime int64       `json:"lastUpdateTime"`
					NextUpdateTime int64       `json:"nextUpdateTime"`
					StatId         interface{} `json:"statId"`
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
				PlanSales           int `json:"planSales"`
				StockSynced         int `json:"stockSynced"`
				SalesPlanDetailList []struct {
					PlanSales    int   `json:"planSales"`
					PlanSendTime int64 `json:"planSendTime"`
				} `json:"salesPlanDetailList"`
			} `json:"salesPlan"`
		} `json:"resultList"`
		ErrorMsgList interface{} `json:"errorMsgList"`
	} `json:"result"`
}

type Login struct {
	Time time.Time
	Qrcode string
}

type Task struct {
	ctx context.Context
	loginC chan *Login
}

func (t *Task) reload() {
	ctx := t.ctx

	listenForNetworkEvent(ctx)
	err := chromedp.Run(ctx,
		network.Enable(),
		t.log("福建爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("福建爬虫结束"),
	)
	if err != nil {
		log.Println("福建爬虫失败:" + err.Error())
	}

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("粤西爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("粤西爬虫选择下拉"),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(2)`),
		t.log("粤西爬虫查询"),
		chromedp.Sleep(time.Second * 5),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("粤西爬虫结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("粤东爬虫开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("粤东爬虫选择下拉"),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(3)`),
		t.log("粤东爬虫查询"),
		chromedp.Sleep(time.Second * 5),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("粤东爬虫结束"),
	)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(time.Second * 10)
}

func (t *Task) log(flag string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		log.Println(flag)
		return nil
	}
}

func (t *Task) screenshotSave(buf *[]byte) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		fileName := fmt.Sprintf("%d.png", time.Now().Nanosecond())
		log.Printf("Write %v", fileName)
		err := ioutil.WriteFile(fileName, *buf, 0644)
		buf = &[]byte{}
		return err
	}
}

func (t *Task) start() {
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

	sel := `#root > div > div > div > main > div > section.login-content.undefined > div > div > div > section > div > div.scan-login.qr-code-activity > div.qr-code > canvas`
	var buf []byte
	err = chromedp.Run(ctx,
		t.log("二维码处理开始"),
		chromedp.WaitEnabled(sel),
		chromedp.Screenshot(sel, &buf, chromedp.NodeVisible),
		t.log("二维码处理结束"),
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

	buf = []byte{}
	err = chromedp.Run(ctx,
		t.log("订单页面开始"),
		chromedp.WaitEnabled(`#mms-header-next > div.mms-header-container > div > div.mms-header__list > a:nth-child(1) > div`),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 3),
		t.log("选择近七日"),
		chromedp.Click(`#dateRange > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > div > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(5)`),
		chromedp.CaptureScreenshot(&buf),
		t.screenshotSave(&buf),
		t.log("订单页面结束"),
	)

	err = chromedp.Run(ctx,
		t.log("选择每页显示数量开始"),
		chromedp.Sleep(time.Second * 2),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div:nth-child(5) > div > div.TB_bottom_1nl4rhj > div > ul > li.PGT_sizeChanger_1nl4rhj > div > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(4)`),
		t.log("选择每页显示数量结束"),
		chromedp.CaptureScreenshot(&buf),
		t.screenshotSave(&buf),
	)
	t.ctx = ctx
	log.Println("初始化结束，进入爬数据阶段...")
}

var requestId network.RequestID

func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*network.EventLoadingFinished); ok {
			if requestId != "" && event.RequestID == requestId {
				go func() {
					var data []byte
					if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
						var err error
						data, err = network.GetResponseBody(requestId).Do(ctx)
						if err != nil {
							return err
						}
						return nil
					})); err != nil {
						fmt.Println("err ", err)
					}

					log.Println("开始导出数据")
					var o OrderResponse
					err := json.Unmarshal(data, &o)
					if err != nil {
						log.Println("导出数据失败2:" + err.Error())
						return
					}
					if o.Result.Total == 0 {
						log.Println("没有可用数据")
						return
					}
					if err = ioutil.WriteFile(string(requestId), data, 0644); err != nil {
						log.Println("导出数据失败3:" + err.Error())
						log.Fatal(err)
					}
					log.Println("结束导出数据")
				}()
			}
		}

		if event, ok := ev.(*network.EventResponseReceived); ok {
			resp := event.Response
			if resp.URL == "https://mc.pinduoduo.com/cartman-mms/orderManagement/pageQueryDetail" {
				requestId = event.RequestID
			}
		}


		//switch ev := ev.(type) {
		//
		//case *network.EventResponseReceived:
		//	resp := ev.Response
		//	if len(resp.Headers) != 0 {
		//		if resp.URL == "https://mc.pinduoduo.com/cartman-mms/orderManagement/pageQueryDetail" {
		//			log.Println(ev.RequestID)
		//			go func (ev *network.EventResponseReceived) {
		//				log.Println("开始导出数据")
		//				c := chromedp.FromContext(ctx)
		//				log.Println(ev.RequestID)
		//				rbp := network.GetResponseBody(ev.RequestID)
		//				log.Println(ev.RequestID)
		//				body, err := rbp.Do(cdp.WithExecutor(context.Background(), c.Target))
		//				if err != nil {
		//					fmt.Println("导出数据失败1:" + err.Error())
		//					return
		//				}
		//				var o OrderResponse
		//				err = json.Unmarshal(body, &o)
		//				if err != nil {
		//					log.Println("导出数据失败2:" + err.Error())
		//					return
		//				}
		//				if o.Result.Total == 0 {
		//					log.Println("没有可用数据")
		//					return
		//				}
		//				if err = ioutil.WriteFile(ev.RequestID.String(), body, 0644); err != nil {
		//					log.Println("导出数据失败3:" + err.Error())
		//					log.Fatal(err)
		//				}
		//			}(ev)
		//		}
		//	}
		//}
		// other needed network Event
	})
}

func (t *Task) Run() {
	t.start()
}

func New(loginC chan *Login, headless bool) *Task {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
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
