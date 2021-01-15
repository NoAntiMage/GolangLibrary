package router

import (
	"fmt"
	"tmpgo/views"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	fmt.Println("router init")
	Router = gin.Default()
}

func init() {
	InitRouter()
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	sysRouter := Router.Group("/sysinfo")
	dbRouter := sysRouter.Group("/db")
	{
		dbRouter.GET("/version", views.GetVersion)
		dbRouter.GET("/now", views.GetNow)
		dbRouter.GET("/dbs", views.GetDatabases)
		dbRouter.GET("/tables", views.GetTables)
	}
}
