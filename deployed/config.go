package deployed

import (
	"log"

	"github.com/spf13/viper"

	"github.com/thinkgos/only-socks5/pkg/infra"
)

var AppConfig = new(Application)

func init() {
	RegisterViperDefault(
		ViperApplicationDefault,
		ViperLoggerDefault,
	)
}

// 载入配置文件
func SetupConfig(path string) {
	err := LoadConfig(path)
	if err != nil {
		log.Fatalf("warning:Parse config file failed: %+v", err)
	}
	// viper.OnConfigChange(func(in fsnotify.Event) {})
	// viper.WatchConfig()

	AppConfig = ViperApplication()
}

// 如果filename为空,将使用config.yaml配置文件,并在当前文件搜索
func LoadConfig(filename string) error {
	// 使用命令行或环境变量给的配置文件,否则使用默认路径进行查找
	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.SetConfigName("config") // 文件名
		viper.SetConfigType("yaml")   // 配置类型
		viper.AddConfigPath(".")      // 增加搜索路径
	}

	ViperInitDefault()

	return viper.ReadInConfig()
}

func IsModeDebug() bool {
	return AppConfig.Mode == infra.ModeDebug
}

func IsModeProd() bool {
	return AppConfig.Mode == infra.ModeProd
}

func IsModeDev() bool {
	return AppConfig.Mode == infra.ModeDev
}
