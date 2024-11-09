package main

import (
	"VizGo/controllers/ChartAPI"
	"VizGo/database"
	"VizGo/jobs/JobChart"
	"VizGo/utils"
)

var logger = utils.GetLogger()

func main() {
	logger.Info("Viz IS Starting")
	database.InitDB()
	JobChart.Run()
	ChartAPI.Run()
}
