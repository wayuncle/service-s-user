package handler

import (
	"github.com/micro/go-micro/v2/web"
	"msp-git.connext.com.cn/connext-go-core/core-util/prouter"
	"service-s-user/handler/userhandler"
)

var User = new(userhandler.User)

func init() {
	User.Version = "v1"
}

func Wrapper() {
	prouter.AddHandlerWrapper(
		prouter.HANDLER_WRAPPER_TIMEELAPSED,
		prouter.HANDLER_WRAPPER_TRACE,
	)
	prouter.AddCallWrapper(
		prouter.CALL_WRAPPER_VALIDATE,
	)
}

// Register 注册、打印路由表
func Register(service web.Service) {
	//Index
	prouter.BindFunction(service, User, User.Index).GET()
	//Create
	prouter.BindFunction(service, User, User.Create).POST()
	//QueryById
	prouter.BindFunction(service, User, User.QueryById).GET()
	//Delete
	prouter.BindFunction(service, User, User.Delete).GET()
	prouter.RouteRegister(service)
}
