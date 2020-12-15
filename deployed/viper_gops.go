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
