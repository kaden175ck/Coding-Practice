package main

import "github.com/cloudwego/hertz/pkg/app"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(200, Response{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}

func FailWithCode(ctx *app.RequestContext, code int, msg string) {
	ctx.JSON(400, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
