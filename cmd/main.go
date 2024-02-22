package main

import (
	"fmt"
	"os"

	"github.com/tankhiwale/ecBackendWeb/config"
	"github.com/tankhiwale/ecBackendWeb/internal/logger"
	"github.com/tankhiwale/ecBackendWeb/internal/server"
)

func main() {

	config, err := config.InitializeConfig("./config", "config")
	if err != nil {
		fmt.Errorf("error initializing config - %v", err)
	}
	if config == nil { //panics if config is nil - null pointer deference
		fmt.Errorf("config cannot be nil")
		os.Exit(1)
	}

	logger := logger.NewLogger()
	configStr := fmt.Sprintf("%#v", config)
	logger.Info(configStr)
	logger.Info("Initializing server..")

	s := server.NewServer(config, logger)
	errC, err := s.Init()

	// to catch pre-server initialization errors
	if err != nil {
		logger.Error("Error: Could not initialize server %s", err)
		os.Exit(1)
	}

	// to handle post-server boot errors
	if err := <-errC; err != nil {
		logger.Error("Error while running server: %s", err)
	}

}
