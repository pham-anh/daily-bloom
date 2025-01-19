package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pham-anh/daily-bloom/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/", controller.Index)
	r.Run(":8080")
}
