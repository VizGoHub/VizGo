package database

import (
	"VizGo/database/models/TChart"
	"VizGo/database/models/TChartDatasets"
	"VizGo/database/models/TDataSource"
	"VizGo/libs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var config, _ = libs.LoadConfig()
var logger = libs.GetLogger()
var DB *gorm.DB

func InitDB() {
	dsn := config.DSN
	var err error
	if config.DB == "mysql" {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if config.DB == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("viz.db"), &gorm.Config{})
	}
	if err != nil {
		logger.Fatal("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&TDataSource.TDataSource{})
	DB.AutoMigrate(&TChart.TChart{})
	DB.AutoMigrate(&TChartDatasets.TChartDatasets{})

}

func GetDB() *gorm.DB {
	return DB
}
