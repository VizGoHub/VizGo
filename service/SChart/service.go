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

func GetChart(chartID int64) TChart.TChart {
	chart := DChart.GetChart(chartID)
	return chart
}

func GetLineChart(chartID int64) map[string]interface{} {
	chart := DChart.GetChart(chartID)
	chartDatasets := DChartDatasets.GetChartDatasets(chartID)
	var datasets []TChartDatasets.JChartDatasets
	var legends []string
	for _, dataset := range chartDatasets {
		item := TChartDatasets.JChartDatasets{
			Name:  dataset.Label,
			Type:  dataset.ChartType,
			Color: dataset.BorderColor,
			LineStyle: TChartDatasets.JLineStyle{
				Width: dataset.BorderWidth,
				Type:  dataset.BorderType,
			},
			Data: toArray(dataset.ColumnData),
		}
		datasets = append(datasets, item)
		legends = append(legends, dataset.Label)
	}

	data := map[string]interface{}{
		"name":     chart.ChartName,
		"labels":   toArray(chart.Labels),
		"datasets": datasets,
		"legends":  legends,
	}

	//sData, _ := json.Marshal(data)
	//logger.Info(string(sss))
	return data
}

func GetBarChart(chartID int64) map[string]interface{} {
	chart := DChart.GetChart(chartID)
	chartDatasets := DChartDatasets.GetChartDatasets(chartID)
	if len(chartDatasets) == 0 {
		return nil
	}

	var datasets = map[string]interface{}{
		"realtimeSort": true,
		"type":         "bar",
		"name":         chartDatasets[0].Label,
		"data":         toArray(chartDatasets[0].ColumnData),
		"label": map[string]interface{}{
			"show":           true,
			"position":       "right",
			"valueAnimation": true,
		},
	}
	var yAxis = map[string]interface{}{
		"type":                    "category",
		"data":                    toArray(chart.Labels),
		"max":                     len(toArray(chart.Labels)),
		"inverse":                 true,
		"animationDuration":       300,
		"animationDurationUpdate": 300,
	}

	data := map[string]interface{}{
		"name":     chart.ChartName,
		"yAxis":    yAxis,
		"datasets": datasets,
	}
	return data
}

func GetTreeMapChart(chartID int64) map[string]interface{} {
	chart := DChart.GetChart(chartID)
	labels := toJsonArray(chart.Labels)
	chartDatasets := DChartDatasets.GetChartDatasets(chartID)
	columnData := toJsonArray(chartDatasets[0].ColumnData)

	var datasets []TChartDatasets.JChartDatasetsTreeMap
	for i, it := range labels {
		item := TChartDatasets.JChartDatasetsTreeMap{
			Name:  it.String(),
			Value: columnData[i].Int(),
		}
		if columnData[i].Int() == 0 {
			item.ItemStyle = TChartDatasets.JItemStyle{
				Color: chartDatasets[0].BorderColor,
			}
		}
		datasets = append(datasets, item)
	}
	if len(datasets) == 0 {
		item := TChartDatasets.JChartDatasetsTreeMap{
			Name:  "No Data",
			Value: 1,
		}
		datasets = append(datasets, item)
	}

	data := map[string]interface{}{
		"type":     "treemap",
		"name":     chart.ChartName,
		"datasets": datasets,
	}
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

func toJsonArray(s string) []gjson.Result {
	return gjson.Parse(s).Array()
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
