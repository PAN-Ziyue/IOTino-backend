package e

import "net/http"

type Status struct {
	Code int
	Msg  string
}

const (
	BadParameter = "非法参数"

	DeviceCreated      = "成功创建设备"
	DeviceNotFound     = "无法找到指定的设备"
	ConflictDevice     = "无法重复创建设备，设备ID或名称重复"
	CannotCreateDevice = "无法创建设备"
	CannotDeleteDevice = "无法删除设备"

	WrongAccount  = "用户名或密码错误"
	DuplicateUser = "用户名或者邮箱已被注册"
	UserNotFound  = "此用户不存在"

	ParseTokenError = "无法验证Token"
	Unauthorized    = "未经鉴权的操作"
	CannotGenToken  = "Token生成失败"
	AuthTimeout     = "Token已超时"
)

func DefaultOk() Status {
	status := Status{}
	status.Code = http.StatusOK
	status.Msg = "ok"
	return status
}

func DefaultError() Status {
	status := Status{}
	status.Code = http.StatusBadRequest
	status.Msg = "error"
	return status
}

func New(Code int, Msg string) Status {
	status := Status{}
	status.Code = Code
	status.Msg = Msg
	return status
}

func (status *Status) Set(Code int, Msg string) {
	status.Code = Code
	status.Msg = Msg
}

func (status *Status) SetCode(Code int) {
	status.Code = Code
}
