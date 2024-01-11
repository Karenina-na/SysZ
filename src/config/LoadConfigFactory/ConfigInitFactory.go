package LoadConfigFactory

import (
	"SysZ/src/exception"
	"SysZ/src/util"
	"errors"
	"github.com/spf13/viper"
)

// InitConfig
// @Description: 初始化配置文件
// @return       E error
func InitConfig() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("InitConfig-config", util.Strval(r))
		}
	}()
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err == nil {
		if err := LoadRootConfig(); err != nil {
			return err
		}
		if err := LoadServerConfig(); err != nil {
			return err
		}
		if err := LoadRoutineConfig(); err != nil {
			return err
		}
	} else {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return exception.NewConfigurationError("InitConfig-config", "配置文件不存在")
		}
	}
	return nil
}
