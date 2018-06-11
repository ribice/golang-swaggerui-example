// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Golang SwaggerUI example
//
// This documentation describes example APIs found under https://github.com/ribice/golang-swaggerui-example
//
//     Schemes: https
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <ribice@gmail.com> https://ribice.ba
//     Host: ribice.ba/goswagg
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
	"github.com/ribice/golang-swaggerui-example/cmd/api"
	_ "github.com/ribice/golang-swaggerui-example/cmd/swagger"
	_ "github.com/ribice/golang-swaggerui-example/cmd/swaggerui" // statik files

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	router.PathPrefix("/swaggerui/").Handler(sh)
	registerV1Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV1Routes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	api.RegisterRepoRoutes(v1, "/repo")
	api.RegisterUserRoutes(v1, "/user")
}
