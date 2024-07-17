package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service-manager/config"
	"service-manager/http/authorization"
	"service-manager/http/common"
)

type AccountController struct {
}

type AccountBody struct {
	Account  string `json:"account,omitempty" form:"account"`
	Password string `json:"password,omitempty" form:"password"`
}

func (controller AccountController) Login(ctx *gin.Context) interface{} {
	var body AccountBody
	if err := ctx.ShouldBind(&body); err != nil {
		panic(common.CodeError{
			Code: -1,
			Err:  err,
		})
	}
	if body.Account == "" || body.Password == "" {
		panic(common.CodeError{
			Code: -1,
			Err:  fmt.Errorf("account or password should not be empty"),
		})
	}
	conf := config.Load()
	pwd, ok := conf.Account[body.Account]
	if !ok {
		panic(common.CodeError{
			Code: -1,
			Err:  fmt.Errorf("account error"),
		})
	}
	if pwd != body.Password {
		panic(common.CodeError{
			Code: -1,
			Err:  fmt.Errorf("password error"),
		})
	}
	token := authorization.NewToken()
	generate, err := token.Generate(body.Account)
	if err != nil {
		panic(common.CodeError{
			Code: -1,
			Err:  err,
		})
	}
	return generate
}

func (controller AccountController) Logout(ctx *gin.Context) interface{} {
	return nil
}

func (controller AccountController) Router(r *gin.RouterGroup) {
	res := common.Response{}
	g := r.Group("")
	g.POST("login", res.SafetyWithData(controller.Login))
	g.GET("logout", res.SafetyWithData(controller.Logout))
}

func NewAccountController() AccountController {
	return AccountController{}
}
