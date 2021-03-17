// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deployed

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/thinkgos/x/builder"
	"github.com/thinkgos/x/extos"
	"github.com/thinkgos/x/lib/habit"

	"github.com/thinkgos/npsocks/pkg/cdir"
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
		configPath, err := cdir.ConfigDir(builder.Name)
		if err != nil {
			return err
		}
		defaultConfigName := "." + builder.Name
		filePath := filepath.Join(configPath, defaultConfigName+".yaml")
		if !extos.IsExist(filePath) {
			os.MkdirAll(configPath, 0755) // nolint: errcheck
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

func IsModeDebug() bool {
	return habit.IsModeDebug(AppConfig.Mode)
}

func IsModeProd() bool {
	return habit.IsModeProd(AppConfig.Mode)
}

func IsModeDev() bool {
	return habit.IsModeDev(AppConfig.Mode)
}
