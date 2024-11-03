package main

import (
	"VizGo/app/api/ChartAPI"
	"VizGo/database"
	"VizGo/jobs/JobChart"
	"VizGo/libs"
)

var logger = libs.GetLogger()

func main() {
	logger.Info("Viz IS Starting")
	database.InitDB()
	JobChart.Run()
	ChartAPI.Run()
}
