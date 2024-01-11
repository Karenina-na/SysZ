package service

import (
	"SysZ/src/config"
	"SysZ/src/exception"
	Factory "SysZ/src/pool"
	"SysZ/src/service/Bean"
	"SysZ/src/util"
	"time"
)

// ServerStatusRoutine
// @Description: 中心状态协程
// @return       E error
func ServerStatusRoutine() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("Register-service", util.Strval(r))
		}
	}()
	util.Loglevel(util.Debug, "ServerStatusRoutine", "创建监控中心协程")
	for {
		select {
		case <-Bean.ServiceCloseChan:
			util.Loglevel(util.Debug, "ServerStatusRoutine", "监控中心协程退出")
			return nil
		case <-time.After(time.Second * time.Duration(config.ListenTime)):
			//获取信息，设置全局信息状态
			Bean.ServerStatus.ServerStatusInfoLock.Lock()
			coreNum, maxNum, activeNum, jobNum := Factory.RoutinePool.CheckStatus()
			Bean.ServerStatus.ServerStatusInfo.SetComputerInfoModel(util.GetCpuInfo(), *util.GetMemInfo(), *util.GetHostInfo(),
				util.GetDiskInfo(), util.GetNetInfo(), coreNum, maxNum, activeNum, jobNum)
			Bean.ServerStatus.ServerStatusInfoLock.Unlock()
		}
	}
}
