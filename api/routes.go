package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hellow worrld"))
	})

	return router
}
