package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const verions = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("\nStarting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
