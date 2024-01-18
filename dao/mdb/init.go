package mdb

import (
	"fmt"
	"sync"
	"time"

	"github.com/twelveeee/amis-admin-go/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mdbOnce sync.Once
var mdbClient *gorm.DB

func GetClient() *gorm.DB {
	return mdbClient
}

func InitMdb() {
	mdbOnce.Do(func() {
		mdbClient = initMySQLClient()
	})
}

func initMySQLClient() *gorm.DB {
	const mysqlDSN string = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	conf := conf.GetAppConf()
	dbDsn := fmt.Sprintf(mysqlDSN, conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dbDsn))
	if err != nil || db == nil {
		for i := 1; i <= 3; i++ {
			db, err = gorm.Open(mysql.Open(dbDsn))
			if db != nil && err == nil {
				break
			}

			time.Sleep(2 * time.Second)
		}

		if err != nil || db == nil {
			panic(err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db
}
