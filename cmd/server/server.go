package server

import (
	"fmt"
	"os"

	"github.com/google/gops/agent"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thinkgos/go-core-package/lib/textcolor"
	"github.com/thinkgos/go-socks5"

	"github.com/thinkgos/only-socks5/deployed"
	"github.com/thinkgos/only-socks5/pkg/builder"
	"github.com/thinkgos/only-socks5/pkg/infra"
	"github.com/thinkgos/only-socks5/pkg/izap"
	"github.com/thinkgos/only-socks5/pkg/sword"
	"github.com/thinkgos/only-socks5/pkg/tip"
)

var configFile string
var port string
var mode string
var Cmd = &cobra.Command{
	Use:          "server",
	Short:        "Start API server",
	Example:      fmt.Sprintf("%s server -c config/config.yaml", builder.Name),
	SilenceUsage: true,
	PreRun:       setup,
	RunE:         run,
	PostRun:      postRun,
}

func init() {
	Cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config/config.yaml", "Start server with provided configuration file")
	Cmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	Cmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,debug,prod")
}

func setup(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags()) // nolint: errcheck
	// viper.SetEnvPrefix("onlys")
	// // OAM_CONFIGFILE
	// viper.BindEnv("config") // nolint: errcheck

	// 1. 读取配置
	deployed.SetupConfig(configFile)
	deployed.SetupLogger()
	infra.WritePidFile()
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println(textcolor.Red("starting server..."))

	go func() {
		infra.HandlerError(agent.Listen(deployed.ViperGops()))
	}()

	showTip()

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(izap.Sugar),
		socks5.WithGPool(sword.AntsPool),
	)

	// Create SOCKS5 proxy on config addr
	return server.ListenAndServe("tcp", deployed.AppConfig.Addr())
}

func postRun(*cobra.Command, []string) {
	infra.RemovePidFile()
}

func showTip() {
	t := tip.Tip{
		textcolor.Red(infra.Banner),
		textcolor.Green(deployed.AppConfig.Name),
		textcolor.Magenta(builder.Version),
		textcolor.Magenta("-h"),
		textcolor.Green("Server run at:"),
		infra.LanIP(),
		deployed.AppConfig.Port,
		textcolor.Green("Server run on PID:"),
		textcolor.Red(cast.ToString(os.Getpid())),
		textcolor.Magenta("Control + C"),
	}
	tip.Show(t)
}
