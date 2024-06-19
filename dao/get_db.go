package dao

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"minimark/config"
	"sync"
)

var (
	db     *gorm.DB
	onceDB sync.Once
)

func GetDB() *gorm.DB {
	onceDB.Do(func() {
		cfg := config.GetConfig()
		driver := cfg.Database.Driver
		dsn := cfg.Database.DSN
		log := logger.Default.LogMode(logger.Info)
		var err error
		switch driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: log,
			})
		case "sqlserver":
			db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
				Logger: log,
			})
		default:
			err = errors.New("不支持的数据库：" + driver)
		}
		if err != nil {
			panic(err)
		}

		// 自定义多对多
		err = db.SetupJoinTable(&Order{}, "Goods", &OrderGoods{})
		if err != nil {
			panic(err)
		}
		err = db.SetupJoinTable(&Procurement{}, "Goods", &ProcurementGoods{})
		if err != nil {
			panic(err)
		}
		// 数据迁移
		err = db.AutoMigrate(&Category{}, &Goods{}, &PayMethod{}, &Order{}, &Procurement{})
		if err != nil {
			panic(err)
		}
	})
	return db
}
