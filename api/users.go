package api

import (
	"context"
	"database/sql"
	"net/http"
)

func (server *Server) getUser(w http.ResponseWriter, r *http.Request) {

	id, err := server.readIDParams(r)

	if err != nil {
		server.notFoundResponse(w, r)
		return
	}

	user, err := server.store.GetUser(context.Background(), id)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			server.notFoundResponse(w, r)
		default:
			server.serverErrorResponse(w, r, err)
		}
		return
	}

	err = server.writeJson(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		server.serverErrorResponse(w, r, err)
	}
}
