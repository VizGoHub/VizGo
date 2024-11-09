package ChartAPI

import (
	"VizGo/models/TChart"
	"VizGo/services/SChart"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Run() {
	router := gin.Default()
	router.Use(Authentication())
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

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Authorization required"})
			c.Abort()
			return
		}

		const prefix = "Basic "
		if !strings.HasPrefix(auth, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Invalid Authorization header"})
			c.Abort()
			return
		}

		// 获取 base64 编码的用户名和密码
		decoded, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Invalid base64 encoding"})
			c.Abort()
			return
		}

		// 分割用户名和密码
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Invalid credentials format"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		// 检查用户名和密码是否匹配
		if username != "admin" || password != "viz" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "Invalid username or password"})
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next()
	}
}
