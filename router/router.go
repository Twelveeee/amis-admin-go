package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller/site"
	"github.com/twelveeee/amis-admin-go/router/middlewares"
)

func Router() {
	// 初始化路由
	router := gin.Default()

	router.Use(middlewares.Cors())

	router.Static("/static", "./templates/static")
	router.LoadHTMLGlob("templates/view/*")

	router.GET("/site", ControllerWrapperForJson(site.NewGetSiteController))

	h5Router(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func h5Router(router *gin.Engine) {
	router.Any("/", ControllerWrapperForHtml("index.html", "Main website"))
}
