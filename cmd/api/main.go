package main

import (
	"net/http"
	"tiagoapi/internal/api"
)

func main() {
	api := &api.Api{Addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.Addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.GetUsersHandler)
	mux.HandleFunc("POST /users", api.CreateUserHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
