package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tankhiwale/ecBackendWeb/config"
	"github.com/tankhiwale/ecBackendWeb/internal/logger"
	"github.com/tankhiwale/ecBackendWeb/internal/server"
)

func main() {
	// TODO: initialize config
	config, err := config.InitializeConfig("./config", "config")
	if err != nil {
		fmt.Errorf("error initializing config - %v", err)
	}
	if config == nil { //panics if config is nil - null pointer deference
		fmt.Errorf("config cannot be nil")
		os.Exit(1)
	}
	// TODO: initialize prometheus
	// TODO: initiliaze logger
	logger := logger.NewLogger()
	configStr := fmt.Sprintf("%#v", config)
	logger.Info(configStr)
	logger.Info("Initializing server..")
	// TODO: initialize server
	s := server.NewServer(config.BindPort)
	s.Init()
	s.Run()

}

type APIFunc func(http.ResponseWriter, *http.Request) error

func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := fn(res, req); err != nil {
			fmt.Errorf("Error : %v", err)
		}
	}
}
