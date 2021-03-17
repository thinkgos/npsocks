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

package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thinkgos/x/builder"

	"github.com/thinkgos/npsocks/pkg/tip"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Get version info",
	Example: fmt.Sprintf("%s version", builder.Name),
	RunE: func(*cobra.Command, []string) error {
		tip.PrintVersion()
		return nil
	},
}
