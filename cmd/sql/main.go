package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mxb/cdd"
	"sort"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/x?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
}

type Cdd struct {
	Id        int
	Date      string
	IsLoad    bool
	From      string
	Data      string
	CreatedAt time.Time
}

type Cdd2 struct {
	Id        int
	Date      string
	IsLoad    bool
	From      string
	Data      string
	CreatedAt time.Time
}

func (c Cdd2) TableName() string {
	return "data2"
}

func main() {
	var dateArr []string
	if err := db.Raw("SELECT `created_date` FROM `data` WHERE `from` = 'order' GROUP BY `created_date` ORDER BY `created_date` ASC").Scan(&dateArr).Error; err != nil {
		panic(err)
	}

	all := map[string][]cdd.OrderResultListResp{}
	for _, date := range dateArr {
		var c []Cdd
		if err := db.Raw("SELECT * FROM `data` WHERE `from` = 'order' AND `created_date` = '" + date + "' ORDER BY `id` ASC").Scan(&c).Error; err != nil {
			panic(err)
		}

		for _, item := range c {
			var o cdd.OrderResp
			if err := json.Unmarshal([]byte(item.Data), &o); err != nil {
				log.Println(item)
				panic(err)
			}
			for _, ii := range o.Result.ResultList {
				d := time.UnixMilli(ii.SessionDate).Format("2006-01-02")
				all[d] = append(all[d], ii)
			}

			keys := make([]string, 0, len(all))
			for k := range all {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, d := range keys {
				a := all[d]
				var o cdd.OrderResp
				o.Success = true
				o.ErrorCode = 1000000
				o.ErrorMsg = ""
				o.Result.Total = len(a)
				o.Result.ResultList = a
				b, err := json.Marshal(o)
				if err != nil {
					panic(err)
				}
				if err := db.Create(&Cdd2{
					Date:   d,
					IsLoad: false,
					From:   "order",
					Data:   string(b),
				}).Error; err != nil {
					panic(err)
				}
			}
		}
	}
}
