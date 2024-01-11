package config

var (
	// Goroutine goroutine参数
	Goroutine struct {
		// MaxRoutineNum goroutine池最大线程数
		MaxRoutineNum int
		// CoreRoutineNum goroutine池核心线程数
		CoreRoutineNum int
		// RoutineTimeOut goroutine池线程超时时间
		RoutineTimeOut int
	}

	// Root root用户信息
	Root struct {
		// RootAccount 超级管理员账号
		RootAccount string
		// RootPassword 超级管理员密码
		RootPassword string
		// TokenEnable 是否开启token验证
		TokenEnable bool
		// TokenExpireTime 超级管理员token过期时间
		TokenExpireTime int
		// TokenSignKey 超级管理员token签名密钥
		TokenSignKey string
	}

	// Port port 端口
	Port string

	// ListenTime 异步更新监控
	ListenTime int64
)
