package e

import "net/http"

type Status struct {
	Code int
	Msg  string
}

const (
	BadJson      = "JSON格式解析失败"
	BadParameter = "非法参数"

	DeviceCreated  = "成功创建设备"
	DeviceNotFound = "无法找到设备"
	ConflictDevice = "无法重复创建设备"

	WrongAccount = "用户名或密码错误"

	ParseTokenError = "无法验证Token"
	CannotGenToken  = "Token生成失败"
	AuthTimeout     = "Token已超时"
)

func DefaultStatus() Status {
	status := Status{}
	status.Code = http.StatusOK
	status.Msg = "ok"
	return status
}

func New(Code int, Msg string) Status {
	status := Status{}
	status.Code = Code
	status.Msg = Msg
	return status
}

func (status Status) Set(Code int, Msg string) {
	status.Code = Code
	status.Msg = Msg
}
