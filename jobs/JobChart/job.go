package JobChart

import (
	"VizGo/database/dao/DChart"
	"VizGo/database/dao/DChartDatasets"
	"VizGo/database/dao/DDataSource"
	"VizGo/database/models/TChart"
	"VizGo/executor/SQLExecutor"
	"VizGo/libs"
	"encoding/json"
	"github.com/jasonlvhit/gocron"
	"time"
)

var logger = libs.GetLogger()
var s = gocron.NewScheduler()

func Run() {
	logger.Info("JobChart Starting")
	s.Every(60).Second().Do(Worker)
	s.Start()
}

func Stop() {
	logger.Info("JobChart Stopping")
	s.Clear()
}

func Worker() {
	tCharts := DChart.GetALLCharts()
	logger.Info("JobChart, tasks=", len(tCharts))
	for _, item := range tCharts {
		time.Sleep(3 * time.Second)
		job(item)
	}
}

func job(chart TChart.TChart) {
	chartDatasets := DChartDatasets.GetChartDatasets(chart.ChartID)

	dataSource := DDataSource.GetDataSource(chart.DataSourceID)

	sqlExecutor := SQLExecutor.SQLExecutor{Database: dataSource.Database, DatabaseUrl: dataSource.DatabaseUrl}
	errInitDB := sqlExecutor.InitDB()
	if errInitDB != nil {
		return
	}

	sqlResults, err := sqlExecutor.Query(chart.QuerySql)

	var labels []string
	if err == nil {
		for _, result := range sqlResults {
			for key := range result {
				if key == "labels" {
					labels = append(labels, result["labels"].(string))
				}
			}
		}
		marshalLabels, errL := json.Marshal(labels)
		if errL == nil {
			DChart.SetLabels(chart.ChartID, string(marshalLabels))
		}

	}

	for _, chartDataset := range chartDatasets {
		var columnData []interface{}
		for _, result := range sqlResults {
			for key := range result {
				if key == chartDataset.ColumnName {
					columnData = append(columnData, result[key])
				}
			}
		}
		if len(columnData) > 0 {
			marshal, errJ := json.Marshal(columnData)
			if errJ == nil {
				DChartDatasets.SetColumnData(chartDataset.ChartDatasetsID, string(marshal))
			}
		}
	}
	sqlExecutor.Free()
}
