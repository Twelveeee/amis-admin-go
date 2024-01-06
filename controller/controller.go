package controller

import (
	"context"
)

type Controller interface {
	Run() error
	CheckParam() error
	Request() interface{}
	Response() interface{}
}

type BaseCtrl struct {
	Ctx  context.Context
	Req  interface{}
	Resp interface{}
}

func (c *BaseCtrl) Run() error            { return nil }
func (c *BaseCtrl) CheckParam() error     { return nil }
func (c *BaseCtrl) Request() interface{}  { return c.Req }
func (c *BaseCtrl) Response() interface{} { return c.Resp }

func NewBaseCtrl(ctx context.Context, req interface{}, resp interface{}) *BaseCtrl {
	return &BaseCtrl{
		Ctx:  ctx,
		Req:  req,
		Resp: resp,
	}
}

type NormalResp struct {
	Errno  int32       `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
