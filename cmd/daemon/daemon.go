package daemon

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takama/daemon"

	"github.com/thinkgos/go-core-package/builder"
)

var srv, _ = daemon.New(builder.Name, "Service", daemon.SystemDaemon)
var CmdInstall = &cobra.Command{
	Use:     "install",
	Short:   "Install the daemon server",
	Example: fmt.Sprintf("%s install", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := srv.Install("server")
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	},
}
var CmdRemove = &cobra.Command{
	Use:     "remove",
	Short:   "Remove the daemon server",
	Example: fmt.Sprintf("%s remove", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := srv.Remove()
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	},
}
var CmdStart = &cobra.Command{
	Use:     "start",
	Short:   "Start the daemon server",
	Example: fmt.Sprintf("%s start", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := srv.Start()
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	},
}
var CmdStop = &cobra.Command{
	Use:     "stop",
	Short:   "Stop the daemon server",
	Example: fmt.Sprintf("%s stop", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := srv.Stop()
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	},
}
var CmdStatus = &cobra.Command{
	Use:     "status",
	Short:   "Get the daemon server status",
	Example: fmt.Sprintf("%s status", builder.Name),
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := srv.Status()
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	},
}
