package TDataSource

type TDataSource struct {
	DataSourceID int64  `gorm:"primaryKey;column:data_source_id" json:"data_source_id"`
	CreateTime   string `gorm:"column:create_time" json:"create_time,omitempty"`
	Database     string `gorm:"column:database" json:"database,omitempty"`
	DatabaseUrl  string `gorm:"column:database_url" json:"database_url,omitempty"`
	Status       int64  `gorm:"column:status" json:"status,omitempty"`
}

func (TDataSource) TableName() string {
	return "t_datasource"
}
