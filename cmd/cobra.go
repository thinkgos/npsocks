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

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thinkgos/go-core-package/builder"
	"github.com/thinkgos/go-core-package/lib/textcolor"

	"github.com/thinkgos/npsocks/cmd/daemon"
	"github.com/thinkgos/npsocks/cmd/version"
)

func init() {
	rootCmd.AddCommand(
		version.Cmd,
		daemon.CmdServer,
		daemon.CmdInstall,
		daemon.CmdRemove,
		daemon.CmdStart,
		daemon.CmdStop,
		daemon.CmdStatus,
	)
}

var rootCmd = &cobra.Command{
	Use:          builder.Name,
	Short:        builder.Name,
	SilenceUsage: true,
	Long:         builder.Name,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip(cmd, args)
			return errors.New(textcolor.Red("requires at least one arg"))
		}
		return nil
	},
	Run: tip,
}

func tip(*cobra.Command, []string) {
	fmt.Printf("欢迎使用 %s %s 可以使用 %s 查看命令\r\n",
		textcolor.Green(builder.Name),
		textcolor.Green(builder.Version),
		textcolor.Red(`-h`))
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

// AddCommand add command
func AddCommand(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}
