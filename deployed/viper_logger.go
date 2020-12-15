package deployed

import (
	"github.com/spf13/viper"
	"github.com/thinkgos/only-socks5/pkg/izap"
)

func ViperLoggerDefault() {
	viper.SetDefault("logger.level", "error")
	viper.SetDefault("logger.console", "console")
	viper.SetDefault("logger.encodeLevel", "LowercaseLevelEncoder")
	viper.SetDefault("logger.writer", "console")
	viper.SetDefault("logger.path", "temp")

	viper.SetDefault("logger.fileName", "onlys.log")
}

func ViperLogger() izap.Config {
	c := viper.Sub("logger")
	return izap.Config{
		Level:       c.GetString("level"),
		Format:      c.GetString("format"),
		EncodeLevel: c.GetString("encodeLevel"),
		Writer:      c.GetString("writer"),
		Stack:       c.GetBool("stack"),
		Path:        c.GetString("path"),

		FileName:   c.GetString("fileName"),
		MaxSize:    c.GetInt("maxSize"),
		MaxAge:     c.GetInt("maxAge"),
		MaxBackups: c.GetInt("maxBackups"),
		LocalTime:  c.GetBool("localTime"),
		Compress:   c.GetBool("compress"),
	}
}
