package SChart

import (
	"VizGo/models/TChart"
	"VizGo/repositories/DChart"
	"VizGo/repositories/DChartDatasets"
	"VizGo/utils"
)

var logger = utils.GetLogger()

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
