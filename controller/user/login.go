package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/twelveeee/amis-admin-go/controller"
	pb "github.com/twelveeee/amis-admin-go/pb_gen/user"
	service "github.com/twelveeee/amis-admin-go/service/user"
)

type loginController struct {
	Ctx  context.Context
	Req  *pb.UserLoginRequest
	Resp *pb.UserLoginResponse
}

func NewLoginController(ctx *gin.Context) controller.Controller {
	return &loginController{
		Ctx:  ctx,
		Req:  &pb.UserLoginRequest{},
		Resp: &pb.UserLoginResponse{},
	}
}

func (c *loginController) Run() error {
	service.Login(c.Ctx, c.Req)
	return nil
}

func (c *loginController) CheckParam() error {
	fmt.Println(c.Req)
	if c.Req.GetUsername() == "" {
		return errors.New("username is empty")
	}
	if c.Req.GetPassword() == "" {
		return errors.New("password is empty")
	}
	return nil
}

func (c *loginController) Request() interface{} {
	return c.Req
}

func (c *loginController) Response() interface{} {
	return c.Resp
}
