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
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/thinkgos/go-core-package/builder"

	"github.com/thinkgos/npsocks/pkg/izap"
)

func ViperLoggerDefault() {
	viper.SetDefault("logger.level", "error")
	viper.SetDefault("logger.console", "console")
	viper.SetDefault("logger.encodeLevel", "LowercaseLevelEncoder")
	viper.SetDefault("logger.writer", "file")
	viper.SetDefault("logger.path", path.Join(os.TempDir(), builder.Name))

	viper.SetDefault("logger.fileName", "npsocks.log")
	viper.SetDefault("logger.maxAge", 3)
}

func ViperLogger() izap.Config {
	c := viper.Sub("logger")
	return izap.Config{
		Level:       c.GetString("level"),
		Format:      c.GetString("format"),
		EncodeLevel: c.GetString("encodeLevel"),
		Writer:      c.GetString("writer"),
		Stack:       c.GetBool("stack"),
		Path:        c.GetString("path"),

		FileName:   c.GetString("fileName"),
		MaxSize:    c.GetInt("maxSize"),
		MaxAge:     c.GetInt("maxAge"),
		MaxBackups: c.GetInt("maxBackups"),
		LocalTime:  c.GetBool("localTime"),
		Compress:   c.GetBool("compress"),
	}
}
