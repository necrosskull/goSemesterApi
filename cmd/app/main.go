package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goSemesterApi/docs"
	"goSemesterApi/internal/controllers"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	api := r.Group("/api")
	{
		api.GET("/week", controllers.GetWeek)
		api.GET("/semester", controllers.GetSemester)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	docs.SwaggerInfo.BasePath = "/api"

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
