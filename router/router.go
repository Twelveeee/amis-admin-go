package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller/site"
	"github.com/twelveeee/amis-admin-go/controller/user"
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
	userRouter(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func h5Router(router *gin.Engine) {
	router.GET("/", ControllerWrapperForHtml("index.html", "Main website"))
}

func userRouter(router *gin.Engine) {
	routerGroup := router.Group("/user")
	routerGroup.Any("/login", ControllerWrapperForJson(user.NewLoginController))
	// routerGroup.GET("/user", ControllerWrapperForJson(site.NewGetSiteController))
	// routerGroup.GET("/createtable", func(c *gin.Context) {
	// 	if err := user.NewUserModel(c).CreateTable(); err != nil {
	// 		c.JSON(200, gin.H{
	// 			"code": 500,
	// 			"msg":  err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 		"msg":  "ok",
	// 	})
	// })
	// routerGroup.GET("/get", func(c *gin.Context) {
	// 	if _, err := user.NewUserModel(c).GetByUserID(1); err != nil {
	// 		c.JSON(200, gin.H{
	// 			"code": 500,
	// 			"msg":  err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 		"msg":  "ok",
	// 	})
	// })
}
