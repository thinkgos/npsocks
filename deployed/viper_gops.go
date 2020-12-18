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

package deployed

import (
	"github.com/google/gops/agent"
	"github.com/spf13/viper"
)

func ViperGops() agent.Options {
	if c := viper.Sub("gops"); c != nil {
		return agent.Options{
			Addr:            c.GetString("addr"),
			ConfigDir:       c.GetString("configDir"),
			ShutdownCleanup: !c.IsSet("cleanup") || c.GetBool("cleanup"),
		}
	}
	return agent.Options{}
}
