package cdir

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"

	"github.com/thinkgos/only-socks5/pkg/infra"
)

const configDirEnvKey = "ONLYS_CONFIG_DIR"

// 配置路径 $HOME/.config/{builder.Name}
func ConfigDir(name string) string {
	if configDir := os.Getenv(configDirEnvKey); configDir != "" {
		return configDir
	}

	if osUserConfigDir, err := os.UserConfigDir(); err == nil && osUserConfigDir != "" {
		return filepath.Join(osUserConfigDir, name)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), name)
	}

	if xdgConfigDir := os.Getenv("XDG_CONFIG_HOME"); xdgConfigDir != "" {
		return filepath.Join(xdgConfigDir, name)
	}

	home, err := homedir.Dir()
	infra.HandlerError(err)
	return filepath.Join(home, ".config", name)
}
