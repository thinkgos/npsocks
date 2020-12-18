package deployed

import (
	"github.com/thinkgos/npsocks/pkg/izap"
)

func SetupLogger() {
	c := ViperLogger()
	logger := izap.New(c)
	izap.ReplaceGlobals(logger)
	izap.Logger.Info("base logger init success")
}
