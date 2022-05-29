package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

type App struct {
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
	a := &App{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux() // If not declared -> DefaultServerMux

	// Routes
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	// Static Files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Starts Server
	listAddr := net.JoinHostPort(addr, port)
	srv := &http.Server{
		Addr:     listAddr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Listening server at %s", listAddr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // Exits if error
}
