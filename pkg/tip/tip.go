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

package tip

import (
	"os"
	"text/template"
)

const tipTpl = `  {{.Banner}}

欢迎使用 {{.Name}} {{.Version}} 可以使用 {{.H}} 查看命令
{{.ServerTitle}}:
	-  Local:   localhost:{{.Port}}
	-  Network: {{.ExtranetIP}}:{{.Port}}
{{.PidTitle}}: {{.PID}}
Enter {{.Kill}} Shutdown Server

`

// Tip 提示信息
type Tip struct {
	Banner      string // 横幅
	Name        string // 应用名称
	Version     string // 应用版本
	H           string // 一般为 -h
	ServerTitle string // 服务标题
	ExtranetIP  string // 外网ip地址
	Port        string // 端口
	PidTitle    string // pid标题
	PID         string // pid
	Kill        string // 一般为 Control + C
}

// Show 显示tip信息到os.Stdout
func Show(t Tip) {
	template.Must(template.New("tip").Parse(tipTpl)).
		Execute(os.Stdout, t) // nolint: errcheck
}
