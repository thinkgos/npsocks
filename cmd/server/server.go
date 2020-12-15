package server

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gops/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thinkgos/go-core-package/lib/textcolor"
	"github.com/thinkgos/go-socks5"

	"github.com/thinkgos/only-socks5/deployed"
	"github.com/thinkgos/only-socks5/pkg/builder"
	"github.com/thinkgos/only-socks5/pkg/infra"
	"github.com/thinkgos/only-socks5/pkg/sword"
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
	// viper.SetEnvPrefix("oam")
	// // OAM_CONFIGFILE
	// viper.BindEnv("config") // nolint: errcheck

	// 1. 读取配置
	deployed.SetupConfig(configFile)

}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println(textcolor.Red("starting server..."))

	go func() {
		infra.HandlerError(agent.Listen(deployed.ViperGops()))
	}()

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithGPool(sword.AntsPool),
	)

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", ":10800"); err != nil {
		panic(err)
	}
	return nil
}

func postRun(*cobra.Command, []string) {

}
