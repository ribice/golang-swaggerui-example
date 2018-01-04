package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterDocRoutes adds swagger endpoint
func RegisterDocRoutes(r *mux.Router, p string) {
	dr := r.PathPrefix(p).Subrouter()
	dr.HandleFunc("", docHandler).Methods("GET")
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	data, _ := ioutil.ReadFile("swagger.json")
	w.Write(data)
}
