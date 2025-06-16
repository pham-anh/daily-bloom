package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pham-anh/daily-bloom/controller"
)

func main() {
	e := echo.New()
	e.Renderer = controller.NewTemplateRenderer("view/*")

	e.GET("/", controller.Number)
	e.POST("/", controller.Number)
	e.GET("/goals/add", controller.GoalsAdd)
	e.GET("/goals/:id", controller.GoalDetail)
	e.Logger.Fatal(e.Start(":8080"))
}
