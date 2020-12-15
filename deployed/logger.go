package deployed

import (
	"github.com/thinkgos/only-socks5/pkg/izap"
)

func SetupLogger() {
	c := ViperLogger()
	logger := izap.New(c)
	izap.ReplaceGlobals(logger)
	izap.Logger.Info("base logger init success")
}
