package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbClient struct {
	Client *gorm.DB
}

var c = &DbClient{}

func init() {
	c = initDbClient()
}

func Client() *gorm.DB {
	return c.Client
}

//to do denpend on config
func initDbClient() *DbClient {

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&timeout=%ds&readTimeout=%ds&writeTimeout=%ds",
		"sz_seller_test",
		"password",
		"host",
		6606,
		"db",
		5,
		5,
		5,
	)
	c, err := gorm.Open(mysql.Open(url))
	if err != nil {
		log.Println("database", "open mysql error "+err.Error())
		panic(err)
	} else {
		log.Println("database connect success ")
	}
	// pool
	sqlDB, err := c.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(600 * time.Second)

	return &DbClient{
		Client: c,
	}
}
