package deployed

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/thinkgos/go-core-package/extos"

	"github.com/thinkgos/only-socks5/pkg/builder"
	"github.com/thinkgos/only-socks5/pkg/infra"
)

var AppConfig = new(Application)

func init() {
	RegisterViperDefault(
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
		configPath := ConfigPath()
		defaultConfigName := "." + builder.Name
		filePath := filepath.Join(configPath, defaultConfigName+".yaml")
		if !extos.IsExist(filePath) {
			if err := extos.WriteFile(filePath, []byte("")); err != nil {
				return err
			}
		}
		viper.AddConfigPath(configPath) // 增加搜索路径
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName) // 文件名
	}

	ViperInitDefault()

	return viper.ReadInConfig()
}

// 配置路径 $HOME/.config/{builder.Name}
func ConfigPath() string {
	home, err := homedir.Dir()
	infra.HandlerError(err)
	return filepath.Join(home, ".config", builder.Name)
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
