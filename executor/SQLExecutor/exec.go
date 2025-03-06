package SQLExecutor

import (
	"VizGo/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time" // 导入 time 包以便设置超时
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

	// 获取底层 *sql.DB 对象
	sqlDB, err := _DB.DB()
	if err != nil {
		return err
	}
	// 设置连接的最大存活时间，超过该时间连接将被自动关闭（此处设置为10分钟）
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	// 可选设置：最大空闲连接数与最大打开连接数
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	ec.DB = _DB
	return nil
}

func (ec *SQLExecutor) Query(querySql string) ([]map[string]interface{}, error) {
	var err error
	var results []map[string]interface{}

	// 执行原生 SQL 查询
	tx := ec.DB.Raw(querySql).Scan(&results)
	if tx != nil {
		return results, err
	}
	return results, err
}

func (ec *SQLExecutor) Free() {
	if ec.DB != nil {
		// 手动关闭数据库连接
		sqlDB, err := ec.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
		ec.DB = nil
	}
}
