package daemon

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takama/daemon"

	"github.com/thinkgos/only-socks5/pkg/builder"
)

var srv, _ = daemon.New(builder.Name, "Service", daemon.SystemDaemon)
var CmdInstall = &cobra.Command{
	Use:     "install",
	Short:   "Install the daemon server",
	Example: fmt.Sprintf("%s install", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := srv.Install("server")
		return err
	},
}
var CmdRemove = &cobra.Command{
	Use:     "remove",
	Short:   "Remove the daemon server",
	Example: fmt.Sprintf("%s remove", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := srv.Remove()
		return err
	},
}
var CmdStart = &cobra.Command{
	Use:     "start",
	Short:   "Start the daemon server",
	Example: fmt.Sprintf("%s start", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := srv.Start()
		return err
	},
}
var CmdStop = &cobra.Command{
	Use:     "stop",
	Short:   "Stop the daemon server",
	Example: fmt.Sprintf("%s stop", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := srv.Stop()
		return err
	},
}
var CmdStatus = &cobra.Command{
	Use:     "status",
	Short:   "Get the daemon server status",
	Example: fmt.Sprintf("%s status", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := srv.Status()
		return err
	},
}
