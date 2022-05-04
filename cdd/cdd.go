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
	"net/http"
	"net/url"
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

type ShouHouResponse struct {
	Success   bool        `json:"success"`
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  interface{} `json:"errorMsg"`
	Result    struct {
		Overview struct {
			SalesTotal          int    `json:"salesTotal"`
			QualityRefundTotal  int    `json:"qualityRefundTotal"`
			QualityRate         string `json:"qualityRate"`
			AfterSaleAmount     int    `json:"afterSaleAmount"`
			PursueSubsidyAmount int    `json:"pursueSubsidyAmount"`
			ReturnedAmount      int    `json:"returnedAmount"`
			ExpectReturnAmount  int    `json:"expectReturnAmount"`
		} `json:"overview"`
		AfterSalesDataPage struct {
			TotalCount int `json:"totalCount"`
			Data       []struct {
				ProductId          int64  `json:"productId"`
				ProductName        string `json:"productName"`
				ImgUrl             string `json:"imgUrl"`
				ProductSpec        string `json:"productSpec"`
				WarehouseGroupId   int    `json:"warehouseGroupId"`
				WarehouseGroupName string `json:"warehouseGroupName"`
				WarehouseDataList  []struct {
					WarehouseId         int    `json:"warehouseId"`
					WarehouseName       string `json:"warehouseName"`
					SalesTotal          int    `json:"salesTotal"`
					AfterSalesTotal     int    `json:"afterSalesTotal"`
					QualityRefundTotal  int    `json:"qualityRefundTotal"`
					AfterSalesRate      string `json:"afterSalesRate"`
					QualityRate         string `json:"qualityRate"`
					AfterSaleAmount     int    `json:"afterSaleAmount"`
					PursueSubsidyAmount int    `json:"pursueSubsidyAmount"`
					ReturnedAmount      int    `json:"returnedAmount"`
					ExpectReturnAmount  int    `json:"expectReturnAmount"`
				} `json:"warehouseDataList"`
			} `json:"data"`
		} `json:"afterSalesDataPage"`
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

func (t *Task) Reload() {
	var err error
	ctx := t.ctx

	err = chromedp.Run(ctx,
		t.log("订单页面初始化开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 5),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 5),
		//移除浮窗
		chromedp.EvaluateAsDevTools("var elemttt = document.getElementById('new_feedback_box_drag_area');elemttt.parentNode.removeChild(elemttt); elemttt = document.getElementById('new_message_box_drag_area');elemttt.parentNode.removeChild(elemttt);", nil),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second * 5),
		t.log("选择近七日"),
		chromedp.Click(`#dateRange > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > div > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(5)`),
		t.log("订单页面初始化结束"),
	)

	err = chromedp.Run(ctx,
		t.log("选择每页显示数量开始"),
		chromedp.Sleep(time.Second * 2),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div:nth-child(5) > div > div.TB_bottom_1nl4rhj > div > ul > li.PGT_sizeChanger_1nl4rhj > div > div > div > div > div > div > div > div.IPT_suffixCell_1nl4rhj.IPT_prefixSuffixCell_1nl4rhj.IPT_pointerCell_1nl4rhj > div > span > i`),
		t.log("选择每页显示数量"),
		chromedp.Sleep(time.Second * 2),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalTopLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(4)`),
		t.log("选择每页显示数量结束"),
	)

	log.Println("初始化结束，进入爬数据阶段...")

	listenForNetworkEvent(ctx)
	err = chromedp.Run(ctx,
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

	err = chromedp.Run(ctx,
		t.log("售后页面初始化开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/after-sales-manage"),
		chromedp.Sleep(time.Second * 3),
		//移除浮窗
		chromedp.EvaluateAsDevTools("var elemttt = document.getElementById('new_feedback_box_drag_area');elemttt.parentNode.removeChild(elemttt); elemttt = document.getElementById('new_message_box_drag_area');elemttt.parentNode.removeChild(elemttt);", nil),
	)

	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second * 3),
		t.log("选择每页30条"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > div.TB_outerWrapper_1nl4rhj.TB_bordered_1nl4rhj.TB_notTreeStriped_1nl4rhj > div.TB_bottom_1nl4rhj > div.TB_bottomRight_1nl4rhj > ul > li.PGT_sizeChanger_1nl4rhj.PGT_alone_1nl4rhj > div > div > div > div > div > div`),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalTopLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(3)`),
		t.log("售后页面初始化结束"),
	)

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("售后福建省开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		t.log("福建省爬虫查询"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("售后福建省结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("售后粤西开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("粤西爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(2)`),
		t.log("粤西爬虫查询"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("售后粤西结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("售后粤东开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("粤东爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(3)`),
		t.log("粤东爬虫查询"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("粤东粤西结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		network.Enable(),
		t.log("售后页面收尾开始"),
		chromedp.Sleep(time.Second * 3),
		t.log("福建省爬虫选择下拉"),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > div:nth-child(1) > div.ST_outerWrapper_1nl4rhj.ST_medium_1nl4rhj > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		t.log("福建省爬虫查询"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div > form > div > button`),
		t.log("售后页面收尾结束"),
	)
	if err != nil {
		log.Println(err)
	}

	err = chromedp.Run(ctx,
		t.log("订单页面收尾开始"),
		chromedp.Sleep(time.Second * 3),
		chromedp.Navigate("https://mc.pinduoduo.com/ddmc-mms/order/management"),
		chromedp.Sleep(time.Second * 3),
		//移除浮窗
		chromedp.EvaluateAsDevTools("var elemttt = document.getElementById('new_feedback_box_drag_area');elemttt.parentNode.removeChild(elemttt); elemttt = document.getElementById('new_message_box_drag_area');elemttt.parentNode.removeChild(elemttt);", nil),
	)

	err = chromedp.Run(ctx,
		t.log("选择福建省"),
		chromedp.Click(`#areaId > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > span > span > div > div > div > div > div > div > div > div.IPT_inputBlockCell_1nl4rhj.ST_inputBlockCell_1nl4rhj`),
		chromedp.Click(`body > div.PT_outerWrapper_1nl4rhj.PP_outerWrapper_1nl4rhj.ST_dropdown_1nl4rhj.ST_mediumDropdown_1nl4rhj.PT_dropdown_1nl4rhj.PT_portalBottomLeft_1nl4rhj.PT_inCustom_1nl4rhj.PP_dropdown_1nl4rhj > div > div > div > div > ul > li:nth-child(1)`),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#root > div.App_mc-content__wmMCn > div.App_mc-main-wrapper__2im7F > main > div > div.management_filter__3Mi1P > form > div > div:nth-child(5) > div.Grid_col_1nl4rhj.Grid_colNotFixed_1nl4rhj.Form_itemWrapper_1nl4rhj > div > div > button:nth-child(1)`),
		t.log("订单页面收尾结束"),
	)

	t.ctx = ctx
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

	err = chromedp.Run(ctx,
		chromedp.WaitEnabled(`#mms-header-next > div.mms-header-container > div > div.mms-header__list > a:nth-child(1) > div`),
	)
	if err != nil {
		log.Println("进入首页失败:" + err.Error())
	}

	t.ctx = ctx
}

var requestId network.RequestID
var urlType string

func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*network.EventLoadingFinished); ok {
			if requestId != "" && event.RequestID == requestId {
				go func() {
					var data []byte
					var err error
					if err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
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
					if urlType == "order" {
						var o OrderResponse
						err = json.Unmarshal(data, &o)
						if err != nil {
							log.Println("导出数据失败2:" + err.Error())
							return
						}
						if o.Result.Total == 0 {
							log.Println("没有可用数据")
							return
						}
					}else {
						var o ShouHouResponse
						err = json.Unmarshal(data, &o)
						if err != nil {
							log.Println("导出数据失败2:" + err.Error())
							return
						}
						if o.Result.AfterSalesDataPage.TotalCount == 0 {
							log.Println("没有可用数据")
							return
						}
					}

					http.PostForm("http://api.mxb.j1mi.com/index/transfer?use=ajax", url.Values{
						"data": {string(data)},
						"from": {urlType},
					})
					//if err = ioutil.WriteFile(string(requestId), data, 0644); err != nil {
					//	log.Println("导出数据失败3:" + err.Error())
					//	log.Fatal(err)
					//}
					log.Println("结束导出数据")
				}()
			}
		}

		if event, ok := ev.(*network.EventResponseReceived); ok {
			resp := event.Response
			if resp.URL == "https://mc.pinduoduo.com/cartman-mms/orderManagement/pageQueryDetail" {
				requestId = event.RequestID
				urlType = "order"
			}
			if resp.URL == "https://mc.pinduoduo.com/ragnaros-mms/after/sales/manage/queryProductAfterSalesStatistic" {
				requestId = event.RequestID
				urlType = "shouhou"
			}
		}
	})
}

func (t *Task) Run() {
	t.start()
}

func New(loginC chan *Login, headless bool) *Task {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
		chromedp.WindowSize(1000, 530),
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
