package main

import (
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"os"
	"os/signal"
	"phone-config/config"
	"phone-config/models"
	"phone-config/web"
	"syscall"
)

var logger *loggo.Logger

func main() {
	conf := config.CollectConfig()

	// Init Logging
	newLogger := loggo.GetLogger("main")
	logger = &newLogger

	err := loggo.ConfigureLoggers("<root>=TRACE")
	if err != nil {
		logger.Errorf("Error configuring Logger: %s", err.Error())
		return
	}
	_, err = loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	if err != nil {
		logger.Errorf("Error configuring Color Logger: %s", err.Error())
		return
	}

	logger.Infof("Starting Phone Config")

	err = models.Init(conf)
	if err != nil {
		logger.Errorf("unable to connect to database: %s", err.Error())
		return
	}

	err = web.Init(conf)
	if err != nil {
		logger.Errorf("unable to start webserver: %s", err.Error())
		return
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof("%s", <-nch)
}
