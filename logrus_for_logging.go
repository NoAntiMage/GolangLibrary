package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	// logger type
	logrus.Debug("Debug")
	logrus.Info("Info")
	logrus.Warn("Warn")
	logrus.Error("Error")
	logrus.Fatal("Fatal")
	logrus.Panic("Panic")

	// logger configuration
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.WithFields(logrus.Fields{
		"name": "wuji",
	}).Info("A name appears")

	// logger output type
	var log = logrus.New()
	file, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"filename": "access",
	}).Info("open file failed")
}