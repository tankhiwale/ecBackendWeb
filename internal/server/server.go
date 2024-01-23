package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	fServer  *fiber.App
	bindPort string
	lock     sync.Mutex
}

func NewServer( /*can take a config struct*/ bindPort string) *server {
	fServer := fiber.New()
	s := &server{
		fServer:  fServer,
		bindPort: bindPort,
	}

	return s
}
func (s *server) Init() error {
	s.fServer.Static("/", "./public")
	s.fServer.Get("/health", checkHealth)
	return nil
}

// BUG: Should end with a log.Fatal(server.Listen()). block with graceful shutdown
func (s *server) Run() error {
	log.Fatal(s.fServer.Listen(fmt.Sprintf(":" + s.bindPort)))
	return nil
}
