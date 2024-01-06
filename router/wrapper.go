package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller"
)

func ControllerWrapperForJson(createFn func(ctx context.Context) controller.Controller) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctrl := createFn(ctx)
		if err := ctrl.CheckParam(); err != nil {
			ctx.JSON(http.StatusInternalServerError, controller.NormalResp{
				Errno:  http.StatusInternalServerError,
				Errmsg: err.Error(),
				Data:   nil,
			})
		}
		if err := ctrl.Run(); err != nil {
			ctx.JSON(http.StatusInternalServerError, controller.NormalResp{
				Errno:  http.StatusInternalServerError,
				Errmsg: err.Error(),
				Data:   nil,
			})
		}

		rep := ctrl.Response()
		ctx.JSON(http.StatusOK, controller.NormalResp{
			Errno:  http.StatusOK,
			Errmsg: "",
			Data:   rep,
		})
	}
}

func ControllerWrapperForHtml(filePath, title string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, filePath, gin.H{
			"title": title,
		})
	}
}

// const amisJsonPTemplate = `(function() {
// 	const response = %s
// 	window.jsonpCallback && window.jsonpCallback(response);
// 	})();`
// func ControllerWrapperForAmisJsonP(createFn func(ctx context.Context) controller.Controller) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctrl := createFn(ctx)
// 		if err := ctrl.CheckParam(); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, controller.NormalResp{
// 				Errno:  http.StatusInternalServerError,
// 				Errmsg: err.Error(),
// 				Data:   nil,
// 			})
// 		}
// 		if err := ctrl.Run(); err != nil {
// 			ctx.JSONP(http.StatusInternalServerError, controller.NormalResp{
// 				Errno:  http.StatusInternalServerError,
// 				Errmsg: err.Error(),
// 				Data:   nil,
// 			})
// 		}

// 		rep := controller.NormalResp{
// 			Errno:  http.StatusOK,
// 			Errmsg: "",
// 			Data:   ctrl.Response(),
// 		}

// 		jsonData, err := json.Marshal(rep)
// 		if err != nil {
// 			ctx.JSONP(http.StatusInternalServerError, controller.NormalResp{
// 				Errno:  http.StatusInternalServerError,
// 				Errmsg: err.Error(),
// 				Data:   nil,
// 			})
// 		}

// 		ctx.String(http.StatusOK, amisJsonPTemplate, string(jsonData))
// 	}
// }
