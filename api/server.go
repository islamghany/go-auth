package api

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	store  string
	router http.Handler
}

func NewServer(s string) *Server {
	return &Server{
		store:  s,
		router: Routes(),
	}

}

func (server *Server) Start(port int) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv.ListenAndServe()
}
