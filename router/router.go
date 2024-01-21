package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller/site"
	"github.com/twelveeee/amis-admin-go/router/middlewares"
)

var authMiddleware *jwt.GinJWTMiddleware

func Router() {
	// 初始化路由
	router := gin.Default()
	var err error

	// router.Use(middlewares.Cors())

	// router.Static("/static", "./templates/static")
	router.LoadHTMLGlob("templates/view/*")

	authMiddleware, err = middlewares.Auth()
	if err != nil {
		panic(err)
	}

	h5Router(router)
	router.POST("/login", authMiddleware.LoginHandler)

	adminRouter(router)

	userRouter(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func h5Router(router *gin.Engine) {
	router.GET("/", ControllerWrapperForHtml("index.html", "Main website"))
	router.GET("/login", ControllerWrapperForHtml("login.html", "Login"))
}

func adminRouter(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	adminRouter.Use(authMiddleware.MiddlewareFunc())

	adminRouter.GET("/site", ControllerWrapperForJson(site.NewGetSiteController))
	adminRouter.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "hello",
		})
	})

}

func userRouter(router *gin.Engine) {
	// routerGroup := router.Group("/user")
	// routerGroup.POST("/login", ControllerWrapperForJson(user.NewLoginController))
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
