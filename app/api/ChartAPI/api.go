package ChartAPI

import (
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
	router.GET("/api/chart/:chartID", func(c *gin.Context) {
		chartID_ := c.Param("chartID")
		chartID, _ := strconv.Atoi(chartID_)
		chart := SChart.GetLineChart(int64(chartID))
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": chart})
	})

	router.Run(":8091")
}
