package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

var port = os.Getenv("PORT")

func main() {
	l := hclog.Default()
	r := mux.NewRouter()
	handlers.CORS()

	r.Handle("/status").Methods("GET")
	r.Handle("/products").Methods("GET")
	r.Handle("/procuts/{slug}/feedback").Methods("POST")

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Info("Starting server...", "port", port)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	l.Info("Received termination signal, shutting down gracefully", "signal", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
