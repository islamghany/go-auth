package api

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
	ScopeAuthorization  = "authorization"
)

func (server *Server) CreateAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := server.readJSON(w, r, &input)

	if err != nil {
		server.badRequestResponse(w, r, err)
		return
	}

	user, err := server.store.GetUserEmail(context.Background(), input.Email)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			server.unauthorizedResponse(w, r)
		default:
			server.serverErrorResponse(w, r, err)
		}
		return
	}

	// Check if the provided password matches the actual password for the user.
	//  match, err := user.Password.Matches(input.Password)
	//  if err != nil {
	// 	 app.serverErrorResponse(w, r, err)
	// 	 return
	//  }
	// If the passwords don't match, then we call the app.invalidCredentialsResponse()
	// helper again and return.
	// if !match {
	//     app.invalidCredentialsResponse(w, r)
	//     return
	// }

	// token, err := utils.GenerateHighEntropyCryptographicallyRandomString(user.ID, 3*24*time.Hour, ScopeActivation)
	// err = server.store.InsertToken(context.Background(), db.InsertTokenParams{
	// 	HashedToken: token.Hash,
	// 	UserID:      token.UserID,
	// 	Expiry:      token.Expiry,
	// 	Scope:       token.Scope,
	// })

	token, err := server.token.CreateToken(user.Email, time.Hour)

	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	// Encode the token to JSON and send it in the response along with a 201 Created
	// status code.
	err = server.writeJson(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		server.serverErrorResponse(w, r, err)
	}
}
