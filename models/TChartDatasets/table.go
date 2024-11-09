package TChartDatasets

type TChartDatasets struct {
	ChartDatasetsID int64  `gorm:"primaryKey;column:chart_datasets_id" json:"chart_datasets_id"`
	CreateTime      int64  `gorm:"column:create_time" json:"create_time,omitempty"`
	ChartID         int64  `gorm:"column:chart_id" json:"chart_id,omitempty"`
	ChartType       string `gorm:"column:chart_type" json:"chart_type,omitempty"`
	ColumnName      string `gorm:"column:column_name" json:"column_name,omitempty"`
	Label           string `gorm:"column:label" json:"label,omitempty"`
	ColumnData      string `gorm:"column:column_data" json:"column_data,omitempty"`
	Status          int64  `gorm:"column:status" json:"status,omitempty"`
}

func (TChartDatasets) TableName() string {
	return "t_chart_datasets"
}
