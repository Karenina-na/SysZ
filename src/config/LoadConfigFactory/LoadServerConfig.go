package LoadConfigFactory

import (
	"SysZ/src/config"
	"SysZ/src/exception"
	"SysZ/src/util"
	"github.com/spf13/viper"
)

// LoadServerConfig
//
//	@Description: 加载服务信息
//	@return E	error
func LoadServerConfig() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("LoadServerConfig-config", util.Strval(r))
		}
	}()
	config.Port = viper.GetString(`server.port`)
	if config.Port == "" {
		config.Port = "8080"
	}
	if !config.VerifyReg(config.PositiveReg, config.Port) {
		return exception.NewConfigurationError("LoadServerConfig-config", "server.Port非法")
	}
	config.ListenTime = int64(viper.GetInt(`server.listen-time`))
	if config.ListenTime == 0 {
		config.ListenTime = 5
	}
	if !config.VerifyReg(config.PositiveReg, util.Strval(config.ListenTime)) {
		return exception.NewConfigurationError("LoadServerConfig-config", "server.listen-time非法")
	}
	return nil
}
