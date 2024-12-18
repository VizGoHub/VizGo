package SQLExecutor

import (
	"VizGo/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = utils.GetLogger()

type SQLExecutor struct {
	Database    string
	DatabaseUrl string
	DB          *gorm.DB
}

func (ec *SQLExecutor) InitDB() error {
	var err error
	var _DB *gorm.DB
	if ec.Database == "mysql" {
		_DB, err = gorm.Open(mysql.Open(ec.DatabaseUrl), &gorm.Config{})
		if err != nil {
			logger.Error("failed to connect to database")
			return fmt.Errorf("failed to connect to database")
		}
	} else if ec.Database == "sqlite" {
		_DB, err = gorm.Open(sqlite.Open(ec.DatabaseUrl), &gorm.Config{})
		if err != nil {
			logger.Error("failed to connect to database")
			return fmt.Errorf("failed to connect to database")
		}
	}

	ec.DB = _DB
	return err
}

func (ec *SQLExecutor) Query(querySql string) ([]map[string]interface{}, error) {
	var err error
	var results []map[string]interface{}

	//logger.Info(querySql)
	tx := ec.DB.Raw(querySql).Scan(&results)
	if tx != nil {
		return results, err
	}
	return results, err
}

func (ec *SQLExecutor) Free() {
	ec.DB = nil
}
