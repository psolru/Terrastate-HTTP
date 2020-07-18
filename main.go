package main

import (
	"github.com/psolru/terrastate-http/config"
	"github.com/psolru/terrastate-http/http"
	"github.com/psolru/terrastate-http/sqlite3"
	"log"
	"sync"
)

func main() {
	log.Println("[MAIN] Starting up...")

	// Create Waitgroup to do startup stuff asynch
	var wg sync.WaitGroup

	// Load Config asynch
	wg.Add(1)
	go func() {
		config.Load()
		defer wg.Done()
	}()

	// Init SQLITE3 asynch
	wg.Add(1)
	go func() {
		sqlite3.Init()
		defer wg.Done()
	}()

	// Start HTTP server when everything is loaded
	wg.Wait()
	http.Listen()
}
