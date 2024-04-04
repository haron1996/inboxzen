package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/haron1996/inboxzen/router"
	"github.com/haron1996/inboxzen/viper"
)

func main() {
	logger := log.New(os.Stderr)

	config, err := viper.LoadConfig(".")
	if err != nil {
		logger.Errorf("Error loading config: %v", err)
	}

	server := &http.Server{
		Addr:     config.PORT,
		Handler:  router.Router(),
		ErrorLog: logger.StandardLog(),
	}

	log.Fatal(server.ListenAndServe())
}
