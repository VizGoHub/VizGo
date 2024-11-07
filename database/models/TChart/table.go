package TChart

type TChart struct {
	ChartID      int64  `gorm:"primaryKey;column:chart_id" json:"chart_id"`
	CreateTime   int64  `gorm:"column:create_time" json:"create_time,omitempty"`
	ChartName    string `gorm:"column:chart_name" json:"chart_name,omitempty"`
	ChartType    string `gorm:"column:chart_type" json:"chart_type,omitempty"`
	Labels       string `gorm:"column:labels" json:"labels,omitempty"`
	DataSourceID int64  `gorm:"column:data_source_id" json:"data_source_id,omitempty"`
	QuerySql     string `gorm:"column:query_sql" json:"query_sql,omitempty"`
	ChartCode    string `gorm:"column:chart_code" json:"chart_code,omitempty"`
	Sort         int64  `gorm:"column:sort" json:"sort,omitempty"`
	Status       int64  `gorm:"column:status" json:"status,omitempty"`
}

func (TChart) TableName() string {
	return "t_chart"
}

type JChart struct {
	ChartID    int64  `json:"chartID"`
	CreateTime string `json:"createTime,omitempty"`
	ChartName  string `json:"chartName,omitempty"`
	ChartType  string `json:"chartType,omitempty"`
	Labels     string `json:"labels,omitempty"`
}
