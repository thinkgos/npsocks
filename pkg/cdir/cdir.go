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

package cdir

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

const configDirEnvKey = "ONLYS_CONFIG_DIR"

// ConfigDir 配置路径
// windows: $APPDATA/{name}
// 类linux: $HOME/.config/{name}
func ConfigDir(name string) (string, error) {
	if configDir := os.Getenv(configDirEnvKey); configDir != "" {
		return configDir, nil
	}

	if osUserConfigDir, err := os.UserConfigDir(); err == nil && osUserConfigDir != "" {
		return filepath.Join(osUserConfigDir, name), nil
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), name), nil
	}

	if xdgConfigDir := os.Getenv("XDG_CONFIG_HOME"); xdgConfigDir != "" {
		return filepath.Join(xdgConfigDir, name), nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", name), nil
}
