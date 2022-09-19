package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (server *Server) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		user := server.contextGetUser(r)

		fmt.Println(user)
		w.Write([]byte("hellow worrld"))
	})
	router.HandlerFunc(http.MethodGet, "/get-user", server.getUser)
	router.HandlerFunc(http.MethodPost, "/token/authenticate/stateful", server.CreateAuthenticationToken)
	router.HandlerFunc(http.MethodPost, "/register", server.registerUser)
	router.HandlerFunc(http.MethodPost, "/token/authenticate/stateless-with-refresh-token", server.CreateAuthenticationTokenWithRenewToken)
	router.HandlerFunc(http.MethodPost, "/token/authenticate/renew-access-token", server.renewAccessToken)
	return server.enableCORs(server.StatelessTokenAuthenticationMiddleware(router))
}
