package jddb

import (
	"encoding/json"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type JdData struct {
	FAuxQty    float64   `json:"f_aux_qty"`
	FDCStockID int       `json:"fdc_stock_id"`
	FItemID    int       `json:"f_item_id"`
	FQty       float64   `json:"f_qty"`
	FSCStockID int       `json:"fsc_stock_id"`
	FInterID   int       `json:"f_inter_id"`
	FDate      time.Time `json:"f_date"`
	FName      string    `json:"f_name"`
	FNumber    string    `json:"f_number"`
	DsName     string    `json:"ds_name"`
	ScName     string    `json:"sc_name"`
}

func Run() {
	query := url.Values{}
	query.Add("database", "AIS20220601165723")
	query.Add("encrypt", "disable")
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "mxb.1234"),
		Host:     "localhost:1433",
		RawQuery: query.Encode(),
	}
	db, err := gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}
	ldate := time.Now().Add(time.Hour * 24 * -7)
	ldateStr := ldate.Format("2006-01-02")
	rows, err := db.Raw("Select e.FAuxQty, e.FDCStockID, e.FItemID, e.FQty, e.FSCStockID, e.FInterID, b.FDate, i.FName, i.FNumber, ds.FName as dsName, sc.FName as scName " +
		"FROM ICStockBill b " +
		"LEFT JOIN ICStockBillEntry e ON b.FInterID = e.FInterID " +
		"LEFT JOIN t_item i ON i.FItemID = e.FItemID " +
		"LEFT JOIN t_Stock ds ON ds.FItemID = e.FDCStockID " +
		"LEFT JOIN t_Stock sc ON sc.FItemID = e.FSCStockID " +
		"WHERE b.FDate > '" + ldateStr + "' " +
		"ORDER BY b.FDate ASC").Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	var dArr []JdData
	for rows.Next() {
		var d JdData
		rows.Scan(&d.FAuxQty, &d.FDCStockID, &d.FItemID, &d.FQty, &d.FSCStockID, &d.FInterID, &d.FDate, &d.FName, &d.FNumber, &d.DsName, &d.ScName)
		dArr = append(dArr, d)
	}
	log.Printf("结束查询，查询到数据:%d", len(dArr))
	buf, err := json.Marshal(dArr)
	if err != nil {
		panic(err)
	}
	log.Println("开始上传数据")
	res, err := http.PostForm("http://api.mxb.j1mi.com/index/sql?use=ajax", url.Values{
		"data": {string(buf)},
	})
	if err != nil {
		log.Println(err)
	}
	resBuf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(resBuf))
	log.Println("结束上传数据")
}
