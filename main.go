package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const (
	DOMAIN = "www.example.org"
	IP = "192.168.100.17"
	PORT = "8080"
)

func newHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	mux.Handle("/", fileServerWithErrors(filepath.Join("public", "static")))
	mux.HandleFunc("/product", handleProduct)

	// avoid slow clients
	return &http.Server{
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
		IdleTimeout: 120*time.Second,
		Handler: mux,
	}
}

func main() {
	// HTTPS server setup
	httpServer := newHTTPServer()
	httpServer.Addr = IP + ":" + PORT

	fmt.Printf("Starting HTTPS server on %s\n", httpServer.Addr)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("httpServer.ListenAndServe failed with %s", err)
	}
}