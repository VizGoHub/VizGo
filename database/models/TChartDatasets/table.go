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
	BorderType      string `gorm:"column:border_type" json:"border_type,omitempty"`
	BorderWidth     int64  `gorm:"column:border_width" json:"border_width,omitempty"`
	ColumnData      string `gorm:"column:column_data" json:"column_data,omitempty"`
	Status          int64  `gorm:"column:status" json:"status,omitempty"`
}

func (TChartDatasets) TableName() string {
	return "t_chart_datasets"
}

type JLineStyle struct {
	Width int64  `json:"width"`
	Type  string `json:"type,omitempty"`
}

type JChartDatasets struct {
	Name      string        `json:"name"`
	Type      string        `json:"type,omitempty"`
	Stack     string        `json:"stack,omitempty"`
	Color     string        `json:"color,omitempty"`
	LineStyle JLineStyle    `json:"lineStyle,omitempty"`
	Data      []interface{} `json:"data,omitempty"`
}
