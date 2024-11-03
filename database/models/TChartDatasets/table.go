package TChartDatasets

type TChartDatasets struct {
	ChartDatasetsID int64  `gorm:"primaryKey;column:chart_datasets_id" json:"chart_datasets_id"`
	CreateTime      string `gorm:"column:create_time" json:"create_time,omitempty"`
	ChartID         int64  `gorm:"column:chart_id" json:"chart_id,omitempty"`
	ChartType       string `gorm:"column:chart_type" json:"chart_type,omitempty"`
	ColumnName      string `gorm:"column:column_name" json:"column_name,omitempty"`
	Label           string `gorm:"column:label" json:"label,omitempty"`
	BackgroundColor string `gorm:"column:background_color" json:"background_color,omitempty"`
	BorderColor     string `gorm:"column:border_color" json:"border_color,omitempty"`
	BorderWidth     int64  `gorm:"column:border_width" json:"border_width,omitempty"`
	ColumnData      string `gorm:"column:column_data" json:"column_data,omitempty"`
	Status          int64  `gorm:"column:status" json:"status,omitempty"`
}

func (TChartDatasets) TableName() string {
	return "t_chart_datasets"
}

type JChartDatasets struct {
	Label           string        `json:"label"`
	BackgroundColor string        `json:"backgroundColor,omitempty"`
	BorderColor     string        `json:"borderColor,omitempty"`
	BorderWidth     int64         `json:"borderWidth,omitempty"`
	Data            []interface{} `json:"data,omitempty"`
}
