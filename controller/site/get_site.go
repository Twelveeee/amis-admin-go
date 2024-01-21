package site

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller"
	"github.com/twelveeee/amis-admin-go/util"
)

type getSiteController struct {
	Ctx  context.Context
	Req  interface{}
	Resp interface{}
}

func NewGetSiteController(ctx *gin.Context) controller.Controller {
	return &getSiteController{
		Ctx:  ctx,
		Req:  nil,
		Resp: nil,
	}
}

func (c *getSiteController) Run() error {
	jsonData, err := util.ReadFileJson("page/site.json")
	if err != nil {
		return err
	}
	c.Resp = jsonData
	return nil
}

func (c *getSiteController) CheckParam() error {
	return nil
}

func (c *getSiteController) Request() interface{} {
	return c.Req
}

func (c *getSiteController) Response() interface{} {
	return c.Resp
}
