package database

import (
	config2 "VizGo/config"
	"VizGo/models/TChart"
	"VizGo/models/TChartDatasets"
	"VizGo/models/TDataSource"
	"VizGo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var config, _ = config2.LoadConfig()
var logger = utils.GetLogger()
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
