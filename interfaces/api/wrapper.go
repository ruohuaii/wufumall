package api

import (
	"context"
	"github.com/ruohuaii/wufumall/infrastructure/capture"

	"github.com/cloudwego/hertz/pkg/app"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 17:14
**/

type Wrapper struct{}

type HandlerFunc func(c context.Context, ctx *app.RequestContext) (any, error)

func (*Wrapper) WrapperHandler(handlerFunc HandlerFunc) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		data, err := handlerFunc(c, ctx)
		ctx.JSON(ctx.Response.StatusCode(), &capture.Response{
			WareError: capture.Handle(err),
			Rlt:       data,
		})
	}
}
