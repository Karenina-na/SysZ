package Bean

import (
	"SysZ/src/entity"
	"sync"
)

// ServerStatusModel is the model of Server status
type ServerStatusModel struct {
	// ServerStatusInfo 服务器状态
	ServerStatusInfo *entity.ComputerInfoModel
	// ServerStatusInfoLock  服务器状态读写锁
	ServerStatusInfoLock sync.RWMutex
}

// NewServerStatusModel
// @Description: 生成ServerStatusModel
// @return       *ServerStatusModel 返回ServerStatusModel
func NewServerStatusModel() *ServerStatusModel {
	return &ServerStatusModel{
		ServerStatusInfo:     entity.NewComputerInfoModel(),
		ServerStatusInfoLock: sync.RWMutex{},
	}
}
