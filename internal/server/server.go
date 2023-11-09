package server

import "sync"

type server struct {
	bindPort string
	lock     sync.Mutex
}

func NewServer( /*can take a config struct*/ bindPort string) *server {

	s := &server{
		bindPort: bindPort,
	}

	return s
}
func (s *server) Init() error {
	return nil
}

func (s *server) Run() error {
	return nil

  // TODO: test
  //NOTE: test
  //BUG:test
  // NOTE:test
  //WARN

}
