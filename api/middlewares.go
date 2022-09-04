package api

import (
	"fmt"
	"net/http"
	"strings"

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
