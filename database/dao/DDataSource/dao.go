package DDataSource

import (
	"VizGo/database"
	"VizGo/database/models/TDataSource"
)

func GetDataSource(dataSourceID int64) TDataSource.TDataSource {
	var row TDataSource.TDataSource
	database.GetDB().
		Model(&TDataSource.TDataSource{}).
		Where(&TDataSource.TDataSource{DataSourceID: dataSourceID}).
		First(&row)
	return row
}
