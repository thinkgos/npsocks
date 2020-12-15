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
