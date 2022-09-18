package api

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	db "github.com/islamghany/go-auth/db/sqlc"
	"github.com/islamghany/go-auth/utils"
	"github.com/tomasen/realip"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
	ScopeAuthorization  = "authorization"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *Server) renewAccessToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := server.readJSON(w, r, &input)
	if err != nil {
		server.badRequestResponse(w, r, err)
		return
	}

	refreshPayload, err := server.token.VerifyToken(input.RefreshToken)
	if err != nil {
		server.unauthorizedResponse(w, r)
		return
	}
	session, err := server.store.GetSession(context.Background(), refreshPayload.ID)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			server.unauthorizedResponse(w, r)
		default:
			server.serverErrorResponse(w, r, err)
		}
		return
	}

	if session.UserID != refreshPayload.UserID {
		server.unauthorizedResponse(w, r)
		return
	}
	if session.RefreshToken != input.RefreshToken {
		server.unauthorizedResponse(w, r)
		return
	}

	if time.Now().After(session.ExpiresAt) {
		server.unauthorizedResponse(w, r)
		return
	}
	accessToken, accessTokenPayLoad, err := server.token.CreateToken(session.UserID, 15*time.Minute)
	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	err = server.writeJson(w, http.StatusCreated, envelope{
		"access_token":            accessToken,
		"access_token_expires_at": accessTokenPayLoad.ExpiredAt,
	}, nil)

	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}
}
func (server *Server) CreateAuthenticationTokenWithRenewToken(w http.ResponseWriter, r *http.Request) {
	input := loginPayload{}

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
	match, err := utils.Matches(input.Password, user.HashedPassword)
	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}
	// If the passwords don't match, then we call the app.invalidCredentialsResponse()
	// helper again and return.
	if !match {
		server.unauthorizedResponse(w, r)
		return
	}

	// creating access token for the user with short life time
	accessToken, accessTokenPayLoad, err := server.token.CreateToken(user.ID, 15*time.Minute)
	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	// creating refresh  token for the user with long life time
	refresh, refreshPayLoad, err := server.token.CreateToken(user.ID, 24*time.Hour*30)
	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	session, err := server.store.InsertSession(context.TODO(), db.InsertSessionParams{
		ID:           refreshPayLoad.ID,
		UserID:       user.ID,
		RefreshToken: refresh,
		ExpiresAt:    refreshPayLoad.ExpiredAt,
		UserAgent:    r.UserAgent(),
		UserIp:       realip.FromRequest(r),
	})

	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}

	/*
			type loginUserResponse struct {
			SessionID             uuid.UUID    `json:"session_id"`
			AccessToken           string       `json:"access_token"`
			AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
			RefreshToken          string       `json:"refresh_token"`
			RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
			User                  db.User 	   `json:"user"`
		}
	*/

	server.setCooke(w, "access_token", accessToken, "/", accessTokenPayLoad.ExpiredAt)
	server.setCooke(w, "refresh_token", refresh, "/token/authenticate/renew-access-token", refreshPayLoad.ExpiredAt)

	err = server.writeJson(w, http.StatusCreated, envelope{
		"user":                     user,
		"session_id":               session.ID,
		"access_token_expires_at":  accessTokenPayLoad.ExpiredAt,
		"refresh_token_expires_at": refreshPayLoad.ExpiredAt,
	}, nil)

	if err != nil {
		server.serverErrorResponse(w, r, err)
		return
	}
}
func (server *Server) CreateAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	input := loginPayload{}

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

	token, _, err := server.token.CreateToken(user.ID, time.Hour)

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
