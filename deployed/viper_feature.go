package deployed

import (
	"github.com/spf13/viper"
	"go.uber.org/atomic"
)

type Feature struct {
	DataScope atomic.Bool // 数据权限功能开关
	OperDB    atomic.Bool // 操作日志写入数据库
	LoginDB   atomic.Bool // 登录日志写入数据库
	OSS       bool        // 使用oss
}

func ViperFeature() *Feature {
	c := viper.Sub("feature")
	ft := &Feature{
		OSS: c.GetBool("oss"),
	}
	ft.DataScope.Store(c.GetBool("dataScope"))
	ft.OperDB.Store(c.GetBool("operDB"))
	ft.LoginDB.Store(c.GetBool("loginDB"))

	return ft
}
