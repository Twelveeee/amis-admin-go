package site

import (
	"context"

	"github.com/twelveeee/amis-admin-go/controller"
	"github.com/twelveeee/amis-admin-go/util"
)

type getJsonpController struct {
	Ctx  context.Context
	Req  interface{}
	Resp interface{}
}

func NewgetJsonpController(ctx context.Context) controller.Controller {
	return &getJsonpController{
		Ctx:  ctx,
		Req:  nil,
		Resp: nil,
	}
}

func (c *getJsonpController) Run() error {
	jsonData, err := util.ReadFileJson("page/jsonp.json")
	if err != nil {
		return err
	}
	c.Resp = jsonData
	return nil
}

func (c *getJsonpController) CheckParam() error {
	return nil
}

func (c *getJsonpController) Request() interface{} {
	return nil
}

func (c *getJsonpController) Response() interface{} {
	return c.Resp
}
