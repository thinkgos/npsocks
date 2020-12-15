package deployed

import (
	"github.com/spf13/viper"

	"github.com/thinkgos/only-socks5/pkg/infra"
)

type Application struct {
	Mode          string // 工作模式
	Name          string // 应用名称
	Host          string // 主机名
	Port          string // 端口
	ReadTimeout   int    // 读超时
	WriterTimeout int    // 写超时
}

func ViperApplicationDefault() {
	viper.SetDefault("mode", infra.ModeProd)
	viper.SetDefault("port", "80")
}

func ViperApplication() *Application {
	return &Application{
		viper.GetString("mode"),
		viper.GetString("name"),
		viper.GetString("host"),
		viper.GetString("port"),
		viper.GetInt("readTimeout"),
		viper.GetInt("writerTimeout"),
	}
}
