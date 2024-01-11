package service

import (
	"SysZ/src/entity"
	"SysZ/src/exception"
	"SysZ/src/service/Bean"
	"SysZ/src/util"
)

// GetServerStatus
// @Description: 获取服务器状态
// @return       C 服务器状态
// @return       E 错误
func GetServerStatus() (C *entity.ComputerInfoModel, E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewUserError("GetServerStatus-service", util.Strval(r))
		}
	}()
	Bean.ServerStatus.ServerStatusInfoLock.RLock()
	defer Bean.ServerStatus.ServerStatusInfoLock.RUnlock()
	return Bean.ServerStatus.ServerStatusInfo, nil
}
