package DChartDatasets

import (
	"VizGo/database"
	"VizGo/models/TChartDatasets"
)

func GetChartDatasets(chartID int64) []TChartDatasets.TChartDatasets {
	var rows []TChartDatasets.TChartDatasets
	database.GetDB().Model(&TChartDatasets.TChartDatasets{}).Where(&TChartDatasets.TChartDatasets{ChartID: chartID, Status: 1}).Find(&rows)
	return rows
}

func SetColumnData(chartDatasetsID int64, columnData string) {
	database.GetDB().
		Model(&TChartDatasets.TChartDatasets{}).
		Where(&TChartDatasets.TChartDatasets{ChartDatasetsID: chartDatasetsID}).
		Updates(TChartDatasets.TChartDatasets{ColumnData: columnData})
}
