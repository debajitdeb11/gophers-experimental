package main

import (
	"log"
	"net/http"
)

type application struct {
	addr string
}

func (s *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hit! - %v", r.RequestURI)
	w.Write([]byte("Hello, World!\n"))
}

func (s *application) handleHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hit! - %v", r.RequestURI)
	w.Write([]byte("Navigating to home page!!!\n"))
}

func (s *application) handleUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hit! - %v", r.RequestURI)
	w.Write([]byte("Navigating to user page!!!\n"))
}

func main() {
	api := &application{
		addr: ":8080",
	}

	// Initilized new mux
	mux := http.NewServeMux()

	svr := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", api.handleHome)
	mux.HandleFunc("GET /user", api.handleUser)

	svr.ListenAndServe()
}
