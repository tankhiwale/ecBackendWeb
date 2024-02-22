package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tankhiwale/ecBackendWeb/config"
	"github.com/tankhiwale/ecBackendWeb/internal/database"
	"github.com/tankhiwale/ecBackendWeb/internal/logger"
	"github.com/tankhiwale/ecBackendWeb/internal/server/handlers"
)

type server struct {
	fServer      *fiber.App
	config       *config.Config
	serverConfig *config.ServerConfig
	logger       logger.ILogger
	DB           *pgxpool.Pool
}

func NewServer(config *config.Config, logger logger.ILogger) *server {

	// middlewares need to be passed on to the fiber server as well
	fiberConfig := fiber.Config{
		ServerHeader: "testServer",
	}
	fServer := fiber.New(fiberConfig)
	s := &server{
		fServer:      fServer,
		config:       config,
		serverConfig: &config.Server,
		logger:       logger,
	}

	return s
}

func (s *server) Init() (<-chan error, error) {

	//Initialize database
	dbConnectionPool, err := database.InitializeDatabase(&s.config.Database, s.logger)
	if err != nil {
		return nil, err
	}

	errC := make(chan error, 1)

	s.DB = dbConnectionPool

	// shutdown logic
	shutDownSignal, shutDownSignalCancel := signal.NotifyContext(context.Background(),
		/*reboot*/ os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT)

	go func() {
		<-shutDownSignal.Done()

		// fiber server shutdown does not reqire a context
		//shutDownTimeout, timeoutCancel = context.WithTimeout(context.Background(), 10*time*time.Second)
		s.logger.Info("Server shutdown signal received")
		defer func() {
			s.logger.Info("closing all doors...")
			//logger sync

			// database pool close
			dbConnectionPool.Close()

			// cancel contexts
			shutDownSignalCancel()
			//timeoutCancel()

			// close channels
			close(errC)

			s.logger.Info("Bye Rukh!")
		}()

		if err := s.fServer.Shutdown(); err != nil {
			errC <- err
		}

	}()

	go s.run(errC, dbConnectionPool)
	return errC, nil
}

// BUG: Should end with a log.Fatal(server.Listen()). block with graceful shutdown
func (s *server) run(errC chan error, repository *pgxpool.Pool) {

	// s.fServer.Static("/", "./public")
	s.fServer.Get("/health", handlers.CheckHealth)
	if err := s.fServer.Listen(fmt.Sprintf(":" + s.serverConfig.BindPort)); err != nil {
		errC <- err
	}
}
