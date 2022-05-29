package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

type app struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Get addr:port from env, default localhost:8080
	port, addr := os.Getenv("PORT"), os.Getenv("LISTEN_ADDR")
	if port == "" {
		port = "8080"
	}
	if addr == "" {
		addr = "localhost"
	}

	// Initialize app
	a := &app{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Starts Server
	listAddr := net.JoinHostPort(addr, port)
	srv := &http.Server{
		Addr:     listAddr,
		ErrorLog: errorLog,
		Handler:  a.routes(),
	}
	infoLog.Printf("Listening server at %s", listAddr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // Exits if error
}
