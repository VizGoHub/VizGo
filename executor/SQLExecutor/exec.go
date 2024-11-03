package SQLExecutor

import (
	"VizGo/libs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = libs.GetLogger()

type SQLExecutor struct {
	Database    string
	DatabaseUrl string
	DB          *gorm.DB
}

func (ec *SQLExecutor) InitDB() {
	var err error
	var _DB *gorm.DB
	if ec.Database == "mysql" {
		_DB, err = gorm.Open(mysql.Open(ec.DatabaseUrl), &gorm.Config{})
	} else if ec.Database == "sqlite" {
		_DB, err = gorm.Open(sqlite.Open(ec.DatabaseUrl), &gorm.Config{})
	}
	if err != nil {
		logger.Fatal("failed to connect to database: %v", err)
	}
	ec.DB = _DB
}

func (ec *SQLExecutor) Query(querySql string) ([]map[string]interface{}, error) {
	var err error
	var results []map[string]interface{}

	logger.Info(querySql)
	tx := ec.DB.Raw(querySql).Scan(&results)
	if tx != nil {
		return results, err
	}
	return results, err
}

func (ec *SQLExecutor) Free() {
	ec.DB = nil
}
