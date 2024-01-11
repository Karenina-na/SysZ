package service

import (
	"SysZ/src/exception"
	Factory "SysZ/src/pool"
	"SysZ/src/service/Bean"
	"SysZ/src/util"
)

// InitServer
// @Description: 初始化服务
// @return       E 异常
func InitServer() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("InitServer-service", util.Strval(r))
		}
	}()
	//设置服务器状态信息异常处理
	util.SetStatusErrorHandle(func(err error) {
		exception.HandleException(exception.NewSystemError("ComputerStatusManager", util.Strval(err)))
	})

	//初始化初始服务器状态存储与协程
	Bean.ServerStatus = Bean.NewServerStatusModel()
	Factory.RoutinePool.CreateWork(ServerStatusRoutine, func(message error) {
		exception.HandleException(message)
	})

	//初始化关闭检测通道
	Bean.ServiceCloseChan = make(chan struct{}, 0)
	return nil
}

// Close
// @Description: 关闭服务
// @return       E 异常
func Close() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("Close-service", util.Strval(r))
		}
	}()
	//关闭检测通道
	close(Bean.ServiceCloseChan)
	return nil
}
