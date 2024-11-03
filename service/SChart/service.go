package SChart

import (
	"VizGo/database/dao/DChart"
	"VizGo/database/dao/DChartDatasets"
	"VizGo/database/models/TChart"
	"VizGo/database/models/TChartDatasets"
	"VizGo/libs"
	"github.com/tidwall/gjson"
)

var logger = libs.GetLogger()

func GetLineChart(chartID int64) map[string]interface{} {
	chart := DChart.GetChart(chartID)
	chartDatasets := DChartDatasets.GetChartDatasets(chartID)
	var datasets []TChartDatasets.JChartDatasets

	for _, dataset := range chartDatasets {
		item := TChartDatasets.JChartDatasets{
			Label:           dataset.Label,
			BackgroundColor: dataset.BackgroundColor,
			BorderColor:     dataset.BorderColor,
			BorderWidth:     dataset.BorderWidth,
			Data:            toArray(dataset.ColumnData),
		}
		datasets = append(datasets, item)
	}

	data := map[string]interface{}{
		"name":     chart.ChartName,
		"labels":   toArray(chart.Labels),
		"datasets": datasets,
	}

	//sData, _ := json.Marshal(data)
	//logger.Info(string(sss))
	return data
}

func toArray(s string) []interface{} {
	dArray := gjson.Parse(s).Array()
	var iArray []interface{}
	for _, it := range dArray {
		iArray = append(iArray, it.Value())
	}
	return iArray
}

func GetALLCharts() []TChart.JChart {
	charts := DChart.GetALLCharts()
	var newCharts []TChart.JChart
	for _, chart := range charts {
		item := TChart.JChart{
			ChartID:    chart.ChartID,
			CreateTime: chart.CreateTime,
			ChartName:  chart.ChartType,
			ChartType:  chart.ChartType,
			Labels:     chart.Labels,
		}
		newCharts = append(newCharts, item)
	}

	return newCharts
}
