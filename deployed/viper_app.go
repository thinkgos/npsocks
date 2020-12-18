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
	"net"

	"github.com/spf13/viper"
)

type Application struct {
	Mode          string // 工作模式
	Name          string // 应用名称
	Host          string // 主机名
	Port          string // 端口
	ReadTimeout   int    // 读超时
	WriterTimeout int    // 写超时
}

func (sf Application) Addr() string {
	return net.JoinHostPort(sf.Host, sf.Port)
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
