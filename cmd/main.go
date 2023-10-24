package main

import (
	"fmt"
	"net/http"
	"os"

	".github.com/Luzik-D/EMTest/internal/api"
	".github.com/Luzik-D/EMTest/internal/storage"
	"github.com/gorilla/mux"
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
	st := storage.New()

	// route
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world")
	}).Methods("GET")

	router.HandleFunc("/api", api.APIHandler(st, logger)).Methods("POST")

	logger.Info("Running server on " + os.Getenv("address"))
	logger.Fatal(http.ListenAndServe(os.Getenv("address"), router))
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
