package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("GET method /"))
			return
		case "/users":
			w.Write([]byte("GET method /users"))
			return
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/users":
			w.Write([]byte("POST method /users"))
			return
		}
	default:
		w.Write([]byte("404 page"))
		return
	}
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	srv.ListenAndServe()
}
