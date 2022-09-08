package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/islamghany/go-auth/db/sqlc"
	"github.com/islamghany/go-auth/token"
)

const JWT_SECRET = "pei3einoh0Beem6uM6Ungohn2heiv5lah1ael4joopie5JaigeikoozaoTew2Eh6"

type envelope map[string]interface{}
type Server struct {
	store *db.Queries
	token token.Maker
}

func NewServer(s *db.Queries) (*Server, error) {

	t, err := token.NewPasetoMaker(JWT_SECRET[:32])
	if err != nil {
		return nil, err
	}
	return &Server{
		store: s,
		token: t,
	}, nil

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
