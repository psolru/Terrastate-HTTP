package http

import (
	"github.com/gorilla/mux"
	"github.com/psolru/terrastate-http/config"
	"github.com/psolru/terrastate-http/env"
	"log"
	"net/http"
	"time"
)

var Port = env.GetEnv("PORT", "8080")

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/list", listHandler).Methods("GET")
	r.HandleFunc("/{ident}", stateHandler).Methods("GET", "POST")
	r.HandleFunc("/{ident}/lock", lockHandler).Methods("LOCK")
	r.HandleFunc("/{ident}/unlock", unlockHandler).Methods("UNLOCK")

	if config.IsAuthActive() {
		r.Use(basicAuth)
	}

	r.Use(accessLog)

	return r
}

// Listen starts the webserver
func Listen() {
	log.Println("[HTTP] Spinning up...")
	srv := &http.Server{
		Handler:      createRouter(),
		Addr:         ":" + Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if config.IsAuthActive() {
		log.Println("[HTTP] BasicAuth is active")
	}

	log.Printf("[HTTP] Listen on port %s...", Port)
	log.Fatal(srv.ListenAndServe())
}
