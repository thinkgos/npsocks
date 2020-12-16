package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	Example:      fmt.Sprintf("%s server -c config.yaml", builder.Name),
	SilenceUsage: true,
	PreRun:       setup,
	RunE:         run,
	PostRun:      postRun,
}

func init() {
	Cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", fmt.Sprintf("config file(default is $HOME/.config/%s/.%s.yaml)", builder.Name, builder.Name))
	Cmd.PersistentFlags().StringVarP(&port, "port", "p", "10800", "Tcp port server listening on")
	Cmd.PersistentFlags().StringVarP(&mode, "mode", "m", "prod", "server mode ; eg:dev,debug,prod")
}

func setup(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags()) // nolint: errcheck
	// viper.SetEnvPrefix("onlys")
	// // ONLYS_CONFIGFILE
	// viper.BindEnv("config") // nolint: errcheck

	// 1. 读取配置
	deployed.SetupConfig(configFile)
	deployed.SetupLogger()
	infra.WritePidFile(deployed.ConfigPath()) // nolint: errcheck
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
	ln, err := net.Listen("tcp", deployed.AppConfig.Addr())
	if err != nil {
		return err
	}
	go server.Serve(ln) // nolint: errcheck

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	fmt.Println(textcolor.Red("close server..."))

	return ln.Close()
}

func postRun(*cobra.Command, []string) {
	infra.RemovePidFile(deployed.ConfigPath()) // nolint: errcheck
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
