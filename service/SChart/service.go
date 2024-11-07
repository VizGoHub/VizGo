package SChart

import (
	"VizGo/database/dao/DChart"
	"VizGo/database/dao/DChartDatasets"
	"VizGo/database/models/TChart"
	"VizGo/libs"
)

var logger = libs.GetLogger()

func GetChartData(chartID int64) map[string]interface{} {
	chart := DChart.GetChart(chartID)
	chartDatasets := DChartDatasets.GetChartDatasets(chartID)
	data := map[string]interface{}{
		"name":     chart.ChartName,
		"chart":    chart,
		"datasets": chartDatasets,
	}
	return data
}

func UpdateChart(chart TChart.TChart) TChart.TChart {
	ret, err := DChart.UpdateChart(chart)
	logger.Info(err)
	return ret
}

func GetALLCharts() []TChart.TChart {
	charts := DChart.GetALLCharts()
	var newCharts []TChart.TChart
	for _, chart := range charts {
		item := TChart.TChart{
			ChartID:    chart.ChartID,
			CreateTime: chart.CreateTime,
			ChartName:  chart.ChartName,
			ChartType:  chart.ChartType,
		}
		newCharts = append(newCharts, item)
	}
	return newCharts
}
