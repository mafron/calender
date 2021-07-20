package main

import (
	"calender/dateControl"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	t := time.Now()
	days := dateControl.GetThisMonthDays(t)

	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "calender.html", days)
	})

	router.Run()
}
