package api

import (
	"context"
	"net/http"

	db "github.com/islamghany/go-auth/db/sqlc"
	"github.com/islamghany/go-auth/utils"
)

var AnonymousUser = &db.User{}

func IsAnonymous(u *db.User) bool {
	return u == AnonymousUser
}

func (server *Server) getUser(w http.ResponseWriter, r *http.Request) {

	user := server.contextGetUser(r)
	err := server.writeJson(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		server.serverErrorResponse(w, r, err)
	}
}

func (server *Server) registerUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := server.readJSON(w, r, &input)
	if err != nil {
		server.badRequestResponse(w, r, err)
		return
	}
	user := db.CreateUserParams{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	hash, err := utils.Set(input.Password)

	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	user.HashedPassword = hash

	u, err := server.store.CreateUser(context.Background(), user)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			server.errorResponse(w, r, http.StatusUnprocessableEntity, "this email is alreay exists")
		default:
			server.serverErrorResponse(w, r, err)
		}
		return
	}

	err = server.writeJson(w, http.StatusCreated, envelope{"user": u}, nil)
	if err != nil {
		server.serverErrorResponse(w, r, err)
	}
}
