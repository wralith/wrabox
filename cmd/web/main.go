package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port, addr := os.Getenv("PORT"), os.Getenv("LISTEN_ADDR")
	if port == "" {
		port = "8080"
	}
	if addr == "" {
		addr = "localhost"
	}

	listAddr := net.JoinHostPort(addr, port)
	log.Printf("listening server at %s", listAddr)

	mux := http.NewServeMux() // If not declared -> DefaultServerMux

	// Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	err := http.ListenAndServe(listAddr, mux)
	log.Fatal(err) // Exits if error
}
