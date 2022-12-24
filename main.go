package main

import (
	"Gmemcached/config"
	"Gmemcached/lib/logger"
	"Gmemcached/tcp"
	EchoHandler "Gmemcached/tcp"
	"fmt"
	"os"
)

const configFile string = "Gmemcached.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "Gmemcached",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})

	if fileExists(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defaultProperties
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{
			Address: fmt.Sprintf("%s:%d",
				config.Properties.Bind,
				config.Properties.Port),
		},
		EchoHandler.MakeHandler())
	if err != nil {
		logger.Error(err)
	}
}
