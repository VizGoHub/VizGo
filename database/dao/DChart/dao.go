package DChart

import (
	"VizGo/database"
	"VizGo/database/models/TChart"
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
