package main

import (
	"calender/dateControl"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	t := time.Now()

	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		days := dateControl.GetThisMonthDays(t)
		ctx.HTML(200, "calender.html", days)
	})

	router.GET("/next", func(ctx *gin.Context) {
		t = t.AddDate(0, 1, 0)
		ctx.Redirect(302, "/")
	})

	router.GET("/back", func(ctx *gin.Context) {
		t = t.AddDate(0, -1, 0)
		ctx.Redirect(302, "/")
	})

	router.GET("/today", func(ctx *gin.Context) {
		t = time.Now()
		ctx.Redirect(302, "/")
	})

	router.Run()
}
