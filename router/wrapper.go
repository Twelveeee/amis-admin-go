package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller"
	"github.com/twelveeee/amis-admin-go/util"
)

type createFunc func(ctx *gin.Context) controller.Controller
type respFunc func(ctx *gin.Context, status, errno int, errmsg string, data interface{})

func requestRunner(createFn createFunc, respFn respFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctrl := createFn(ctx)

		req := ctrl.Request()

		if err := ctx.ShouldBind(&req); err != nil {
			respFn(ctx, http.StatusInternalServerError, util.SystemErr, err.Error(), nil)
			return
		}

		// if req == nil {
		// 	respFn(ctx, http.StatusInternalServerError, util.SystemErr, "req nil", nil)
		// 	return
		// }

		if err := ctrl.CheckParam(); err != nil {
			respFn(ctx, http.StatusOK, util.ParamErr, err.Error(), nil)
			return
		}
		if err := ctrl.Run(); err != nil {
			respFn(ctx, http.StatusInternalServerError, util.SystemErr, err.Error(), nil)
			return
		}

		rep := ctrl.Response()
		respFn(ctx, http.StatusOK, util.Success, "", rep)
	}
}

func ControllerWrapperForJson(createFn createFunc) gin.HandlerFunc {
	return requestRunner(createFn, jsonResponse)
}

func jsonResponse(ctx *gin.Context, status, errno int, errmsg string, data interface{}) {
	resp := map[string]interface{}{
		"errno":  errno,
		"errmsg": errmsg,
		"data":   data,
	}
	ctx.JSON(status, resp)
}

func ControllerWrapperForHtml(filePath, title string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		host := ctx.Request.Host
		// 获取协议
		protocol := "http://"
		if ctx.Request.TLS != nil {
			protocol = "https://"
		}
		host = protocol + host
		ctx.HTML(http.StatusOK, filePath, gin.H{
			"title":    title,
			"API_HOST": host,
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
