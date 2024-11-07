package ChartAPI

import (
	"VizGo/database/models/TChart"
	"VizGo/service/SChart"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Run() {
	router := gin.Default()

	router.Static("/static", "./dist/static")
	router.StaticFile("/", "./dist/index.html")

	router.GET("/api/charts", func(c *gin.Context) {
		charts := SChart.GetALLCharts()
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": charts})
	})

	router.GET("/api/chartData/:chartID", func(c *gin.Context) {
		chartID_ := c.Param("chartID")
		chartID, _ := strconv.Atoi(chartID_)
		chartData := SChart.GetChartData(int64(chartID))
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": chartData})
	})

	router.POST("/api/updateChart", func(c *gin.Context) {
		var chart TChart.TChart
		if err := c.BindJSON(&chart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Invalid request data"})
			return
		}
		chartData := SChart.UpdateChart(chart)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": chartData})
	})

	router.Run(":8091")
}
