package DChart

import (
	"VizGo/database"
	"VizGo/models/TChart"
)

func GetALLCharts() []TChart.TChart {
	var rows []TChart.TChart
	database.GetDB().Model(&TChart.TChart{}).Where(&TChart.TChart{Status: 1}).Order(" sort DESC ").Find(&rows)
	return rows
}

func SetLabels(chartID int64, labels string) {
	database.GetDB().
		Model(&TChart.TChart{}).
		Where(&TChart.TChart{ChartID: chartID}).
		Updates(TChart.TChart{Labels: labels})
}

func GetChart(chartID int64) TChart.TChart {
	var row TChart.TChart
	database.GetDB().Model(&TChart.TChart{}).Where(&TChart.TChart{ChartID: chartID, Status: 1}).First(&row)
	return row
}

func UpdateChart(chart TChart.TChart) (TChart.TChart, error) {
	var row = TChart.TChart{
		ChartName:    chart.ChartName,
		ChartType:    chart.ChartType,
		DataSourceID: chart.DataSourceID,
		QuerySql:     chart.QuerySql,
		ChartCode:    chart.ChartCode,
	}
	var err error
	if chart.ChartID == 0 {
		err = database.GetDB().Model(&TChart.TChart{}).Create(&row).Error
	} else {
		err = database.GetDB().
			Model(&TChart.TChart{}).
			Where(&TChart.TChart{ChartID: chart.ChartID}).
			Updates(&row).Error
	}
	return row, err
}
