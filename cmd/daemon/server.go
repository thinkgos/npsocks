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

package daemon

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/gops/agent"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thinkgos/go-core-package/extos"
	"github.com/thinkgos/go-core-package/lib/habit"
	"github.com/thinkgos/go-core-package/lib/textcolor"
	"github.com/thinkgos/go-socks5"

	"github.com/thinkgos/go-core-package/builder"

	"github.com/thinkgos/npsocks/deployed"
	"github.com/thinkgos/npsocks/pkg/cdir"
	"github.com/thinkgos/npsocks/pkg/infra"
	"github.com/thinkgos/npsocks/pkg/izap"
	"github.com/thinkgos/npsocks/pkg/tip"
)

var configFile string
var port string
var mode string
var CmdServer = &cobra.Command{
	Use:          "server",
	Short:        "Start API server",
	Example:      fmt.Sprintf("%s server -c config.yaml", builder.Name),
	SilenceUsage: true,
	PreRun:       setup,
	RunE:         run,
	PostRun:      postRun,
}

func init() {
	CmdServer.PersistentFlags().StringVarP(&configFile, "config", "c", "", fmt.Sprintf("config file(default is $HOME/.config/%s/.%s.yaml)", builder.Name, builder.Name))
	CmdServer.PersistentFlags().StringVarP(&port, "port", "p", "10800", "Tcp port server listening on")
	CmdServer.PersistentFlags().StringVarP(&mode, "mode", "m", "prod", "server mode ; eg:dev,debug,prod")
}

func setup(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags()) // nolint: errcheck
	// viper.SetEnvPrefix("npsocks")
	// // NPSOCKS_CONFIGFILE
	// viper.BindEnv("config") // nolint: errcheck

	// 1. 读取配置
	deployed.SetupConfig(configFile)
	deployed.SetupLogger()
	configDir, err := cdir.ConfigDir(builder.Name)
	if err != nil {
		log.Fatal(err)
	}
	if err = habit.WritePidFile(configDir); err != nil {
		log.Fatal(err)
	}
}

func run(*cobra.Command, []string) (err error) {
	exec := &Executable{}
	if extos.IsWindows() {
		_, err = srv.Run(exec)
	} else {
		exec.Run()
	}
	return
}

func postRun(*cobra.Command, []string) {
	agent.Close()
	configDir, err := cdir.ConfigDir(builder.Name)
	if err != nil {
		log.Println(err)
		return
	}
	habit.RemovePidFile(configDir) // nolint: errcheck
}

func showTip() {
	t := tip.Tip{
		Banner:      textcolor.Red(infra.Banner),
		Name:        textcolor.Green(deployed.AppConfig.Name),
		Version:     textcolor.Magenta(builder.Version),
		H:           textcolor.Magenta("-h"),
		ServerTitle: textcolor.Green("Server run at:"),
		ExtranetIP:  infra.LanIP(),
		Port:        deployed.AppConfig.Port,
		PidTitle:    textcolor.Green("Server run on PID:"),
		PID:         textcolor.Red(cast.ToString(os.Getpid())),
		Kill:        textcolor.Magenta("Control + C"),
	}
	tip.Show(t)
}

type Executable struct{}

func (Executable) Start() {}
func (Executable) Stop()  {}
func (Executable) Run() {
	fmt.Println(textcolor.Red("starting server..."))

	go func() {
		infra.HandlerError(agent.Listen(deployed.ViperGops()))
	}()

	showTip()

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(izap.Sugar),
		socks5.WithGPool(deployed.AntsPool),
	)
	ln, err := net.Listen("tcp", deployed.AppConfig.Addr())
	if err != nil {
		fmt.Println(err)
		return
	}
	go server.Serve(ln) // nolint: errcheck

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	fmt.Println(textcolor.Red("close server..."))
	ln.Close()
}
