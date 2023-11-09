package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tankhiwale/ecBackendWeb/internal/server"
)

func main() {
  //initialize config

  //initialize server
  s := server.NewServer(":27200")
  s.Init()
  s.Run()

  
}

type APIFunc func(http.ResponseWriter, *http.Request) error

func makeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := fn(res, req); err != nil {
			fmt.Errorf("Error : " + err.Error())
		}
	}
}

