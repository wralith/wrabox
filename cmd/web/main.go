package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wralith/wrabox/pkg/models/mysql"
)

type app struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	dsn := flag.String("dsn", "web:WebPass1!@/wrabox?parseTime=true", "MYSQL DSN")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	db.Ping()

	defer db.Close()

	// Get addr:port from env, default localhost:8080
	port, addr := os.Getenv("PORT"), os.Getenv("LISTEN_ADDR")
	if port == "" {
		port = "8080"
	}
	if addr == "" {
		addr = "localhost"
	}

	templateCache, err := newTemplateCache("./web/template/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize app
	a := &app{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Starts Server
	listAddr := net.JoinHostPort(addr, port)
	srv := &http.Server{
		Addr:     listAddr,
		ErrorLog: errorLog,
		Handler:  a.routes(),
	}
	infoLog.Printf("Listening server at %s", listAddr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err) // Exits if error
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
