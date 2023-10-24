package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	// load environment
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// init logger
	logger := initLogger(os.Getenv("env"))
	logger.Debug("debug mode")
	logger.Info("info mode")

	// init storage

	// route

	// run server

}

func initLogger(env string) *logrus.Logger {
	log := logrus.New()

	switch env {
	case "dev":
		log.SetOutput(os.Stdout)
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.TextFormatter{})
	case "debug":
		log.SetOutput(os.Stderr)
		log.SetLevel(logrus.DebugLevel)
		log.SetFormatter(&logrus.TextFormatter{})
	default:
		log.SetOutput(os.Stdout)
		log.SetLevel(logrus.InfoLevel)
	}

	return log

}
