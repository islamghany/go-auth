package api

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/islamghany/go-auth/db/sqlc"
	"github.com/islamghany/go-auth/utils"
)

func (server *Server) BasicAuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			server.forbiddenResponse(w, r)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Basic" {
			server.unauthorizedResponse(w, r)
			return
		}
		cred, err := utils.DecodeBasicAuthBase(headerParts[1])
		if err != nil {
			server.forbiddenResponse(w, r)
			return
		}
		fmt.Println(cred)
		next.ServeHTTP(w, r)
	})
}

func (server *Server) StatefulTokenAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Add the "Vary: Authorization" header to the response. This indicates to any
		// caches that the response may vary based on the value of the Authorization
		// header in the request.
		w.Header().Add("Vary", "Authorization")

		// Retrieve the value of the Authorization header from the request. This will
		// return the empty string "" if there is no such header found.
		authorizationHeader := r.Header.Get("Authorization")

		// If there is no Authorization header found, use the contextSetUser() helper
		// that we just made to add the AnonymousUser to the request context. Then we
		// call the next handler in the chain and return without executing any of the
		// code below.
		if authorizationHeader == "" {
			r = server.contextSetUser(r, AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise, we expect the value of the Authorization header to be in the format
		// "Bearer <token>". We try to split this into its constituent parts, and if the
		// header isn't in the expected format we return a 401 Unauthorized response
		// using the invalidAuthenticationTokenResponse() helper (which we will create
		// in a moment).
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			server.unauthorizedResponse(w, r)
			return
		}

		// Extract the actual authentication token from the header parts.
		token := headerParts[1]

		// Retrieve the details of the user associated with the authentication token,
		// again calling the invalidAuthenticationTokenResponse() helper if no
		// matching record was found. IMPORTANT: Notice that we are using
		// ScopeAuthentication as the first parameter here.
		tokenHash := sha256.Sum256([]byte(token))
		fmt.Println(tokenHash)

		user, err := server.store.GetUserFromToken(context.Background(), db.GetUserFromTokenParams{
			Scope:       ScopeAuthentication,
			HashedToken: tokenHash[:],
			Expiry:      time.Now(),
		})
		fmt.Println(err)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				server.unauthorizedResponse(w, r)
			default:
				server.serverErrorResponse(w, r, err)
			}
			return
		}

		// Call the contextSetUser() helper to add the user information to the request
		// context.
		r = server.contextSetUser(r, &user)

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func (server *Server) StatelessTokenAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Add the "Vary: Authorization" header to the response. This indicates to any
		// caches that the response may vary based on the value of the Authorization
		// header in the request.
		w.Header().Add("Vary", "Authorization")

		// Retrieve the value of the Authorization header from the request. This will
		// return the empty string "" if there is no such header found.
		authorizationHeader := r.Header.Get("Authorization")

		// If there is no Authorization header found, use the contextSetUser() helper
		// that we just made to add the AnonymousUser to the request context. Then we
		// call the next handler in the chain and return without executing any of the
		// code below.
		if authorizationHeader == "" {
			r = server.contextSetUser(r, AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise, we expect the value of the Authorization header to be in the format
		// "Bearer <token>". We try to split this into its constituent parts, and if the
		// header isn't in the expected format we return a 401 Unauthorized response
		// using the invalidAuthenticationTokenResponse() helper (which we will create
		// in a moment).
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			server.unauthorizedResponse(w, r)
			return
		}

		// Extract the actual authentication token from the header parts.
		token := headerParts[1]

		p, err := server.token.VerifyToken(token)

		if err != nil {
			server.unauthorizedResponse(w, r)
			return
		}
		fmt.Println(p)
		user, err := server.store.GetUser(context.Background(), p.UserID)
		fmt.Println(err)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				server.unauthorizedResponse(w, r)
			default:
				server.serverErrorResponse(w, r, err)
			}
			return
		}

		// Call the contextSetUser() helper to add the user information to the request
		// context.
		r = server.contextSetUser(r, &user)
		next.ServeHTTP(w, r)
	})
}
