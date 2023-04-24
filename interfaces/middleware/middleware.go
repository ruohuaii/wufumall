package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ruohuaii/wufumall/domain/entity"
	"github.com/ruohuaii/wufumall/domain/repository"
	"github.com/ruohuaii/wufumall/infrastructure/capture"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:38
**/

type Handler func(c context.Context, ctx *app.RequestContext) (any, error)

func Wrap(handler Handler) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		data, err := handler(c, ctx)
		ctx.JSON(ctx.Response.StatusCode(), &capture.Response{
			WareError: capture.Handle(err),
			Rlt:       data,
		})
	}
}

//RequestTrackingMiddleware 追踪请求
func RequestTrackingMiddleware(repo repository.RequestRecordRepository) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		requestData := string(ctx.Request.Body())
		requestID := ctx.Response.Header.Get("X-Request-ID")
		var userid uint64 = 1
		requestIP := ctx.ClientIP()
		requestURI := string(ctx.Request.RequestURI())
		err := repo.Save(&entity.RequestRecord{
			UserId:      userid,
			RequestId:   requestID,
			RequestIP:   requestIP,
			RequestURI:  requestURI,
			RequestData: requestData,
		})
		if err != nil {
			ctx.JSON(502, "系统异常!")
			return
		}
		ctx.Next(c)
	}
}
