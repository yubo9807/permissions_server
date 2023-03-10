package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

type StateType struct {
	Code    int
	Data    any
	Message string
	RunTime string
}

var State = StateType{}

var startTime time.Time

// 初始化
func (state *StateType) Init() {
	startTime = time.Now()
	state.Code = 400
	state.Message = "unknown error"
	state.RunTime = startTime.String()
	state.Data = nil
}

// 返回统一格式
func (state *StateType) Result(ctx *gin.Context) {
	state.RunTime = time.Since(startTime).String()
	ctx.JSON(200, gin.H{
		"code":    state.Code,
		"data":    state.Data,
		"message": state.Message,
		"runTime": state.RunTime,
	})
}

// 参数错误
func (state *StateType) ErrorParams() {
	state.Code = 406
	state.Message = "params error"
}

// 未授权
func (state *StateType) ErrorUnauthorized() {
	state.Code = 401
	state.Message = "unauthorized"
}

// 自定义错误消息
func (state *StateType) ErrorCustom(msg string) {
	state.Code = 500
	state.Message = msg
}

// 请求成功，并返回数据
func (state *StateType) SuccessData(data interface{}) {
	state.Code = 200
	state.Message = "success"
	state.Data = data
}

// 请求成功
func (state *StateType) Success() {
	state.SuccessData("success")
}

func (state *StateType) GetData() interface{} {
	return state.Data
}
