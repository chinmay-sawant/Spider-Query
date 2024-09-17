package internal

import (
	"chinmay-sawant/Spider-Query/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbs model.Dbs) *gorm.DB {
	// dsn := "host=localhost user=sc password=sc dbname=sc port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := dbs.GetDbString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Exec("SET search_path TO sc")
	if err != nil {
		fmt.Println("Unable to connect to DB")
		return nil
	}
	return db
}
