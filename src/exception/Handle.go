package exception

import (
	"SysZ/src/util"
	"os"
)

// HandleException
// @Description: Handle the exception
// @param        err : The exception
func HandleException(err interface{}) {
	switch E := err.(type) {
	case *ConfigurationError:
		configurationExHandle(E)
	case *SystemError:
		systemExHandle(E)
	case *UserError:
		userExHandle(E)
	default:
		util.Loglevel(util.Error, "未知错误", util.Strval(err))
		os.Exit(0)
	}
}

// configurationExHandle
// @Description: Handle the configuration exception
// @param        err : The exception
func configurationExHandle(err *ConfigurationError) {
	util.Loglevel(util.Warn, err.Name, err.Message)
	os.Exit(0)
}

// systemExHandle
// @Description: Handle the system exception
// @param        err : The exception
func systemExHandle(err *SystemError) {
	util.Loglevel(util.Warn, err.Name, err.Message)
	os.Exit(0)
}

// userExHandle
// @Description: Handle the user exception
// @param        err : The exception
func userExHandle(err *UserError) {
	util.Loglevel(util.Warn, err.Name, err.Message)
}
