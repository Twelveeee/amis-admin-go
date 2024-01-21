package middlewares

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	pb "github.com/twelveeee/amis-admin-go/pb_gen/user"
	service "github.com/twelveeee/amis-admin-go/service/user"
)

const identityKey = "_aik"

func Auth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "xxxxsssswwww",
		SigningAlgorithm: "HS512",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc:      payloadFunc,
		IdentityHandler:  identityHandler,
		Authenticator:    authenticator,
		Authorizator:     authorizator,
		Unauthorized:     unauthorized,
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,

		SendCookie:     true,
		SecureCookie:   false,
		CookieHTTPOnly: true,
		// CookieDomain:   "192.168.102.132:8080",
		CookieName:     "jwt",
		TokenLookup:    "header: Authorization, query: token, cookie: jwt",
		CookieSameSite: http.SameSiteNoneMode,
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*pb.UserLoginResponse); ok {
		return jwt.MapClaims{
			identityKey: v.UserId,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &pb.UserLoginResponse{
		UserId: claims[identityKey].(string),
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	loginReq := &pb.UserLoginRequest{}
	loginResp := &pb.UserLoginResponse{}

	if err := c.ShouldBind(loginReq); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	fmt.Println("req username:", loginReq.Username, loginReq.Password)
	loginResp, err := service.Login(c, loginReq)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return loginResp, nil
}

func authorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*pb.UserLoginResponse); ok {
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"errno":  code,
		"errmsg": message,
	})
}
