package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"net/url"
	"time"
)

func main() {
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
	rows, err := db.Raw("Select Top 5 FDate FROM ICStockBill ORDER BY FInterID DESC").Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var date time.Time
		if err := rows.Scan(&date); err != nil {
			log.Println(err)
		}
		log.Println(date)
	}
}
