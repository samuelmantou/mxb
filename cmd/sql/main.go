package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := "sqlserver://sa:mxb.1234@localhost:1433?database=AIS20220601165723"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}
}
