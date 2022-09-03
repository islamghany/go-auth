package api

import (
	"fmt"
	"net/http"
)

func (server *Server) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := server.writeJson(w, status, env, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}
func (server *Server) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//server.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	server.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (server *Server) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	server.errorResponse(w, r, http.StatusNotFound, message)
}
