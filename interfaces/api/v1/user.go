package v1

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ruohuaii/wufumall/application"
	model "github.com/ruohuaii/wufumall/interfaces/model/user"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:20
**/

type User struct {
	userApp application.UserAppInterface
}

func NewUser(userApp application.UserAppInterface) *User {
	return &User{userApp: userApp}
}

func (u *User) SignUp(c context.Context, ctx *app.RequestContext) (any, error) {
	request := model.SignUpRequest{}
	err := ctx.BindAndValidate(&request)
	if err == nil {
		return nil, errors.New("这是我定义的错误")
	}

	type Output struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	return &Output{Name: "hhs", Age: 18}, nil
}

func (u *User) SignIn() {
	
}
