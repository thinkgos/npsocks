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
