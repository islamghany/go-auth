package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/islamghany/go-auth/db/sqlc"
)

type envelope map[string]interface{}
type Server struct {
	store *db.Queries
}

func NewServer(s *db.Queries) *Server {
	return &Server{
		store: s,
	}

}

func (server *Server) Start(port int) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv.ListenAndServe()
}
