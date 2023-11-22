package database

import (
	"acs/src/models"
	"acs/src/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassWord,
		utils.DbName,
		utils.DbPort,
		utils.SslMode,
	)
	//使用 postgresql 的driver 连接数据库。
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic(err)
	}

	DB = db

	//自动迁移类到数据库。
	err = db.AutoMigrate(&models.Car{}, &models.Location{}, &models.Customer{}, &models.BookRecord{})
	if err != nil {
		panic("Could not create a table")
	}
	initData()
}
