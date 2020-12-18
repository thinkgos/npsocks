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
