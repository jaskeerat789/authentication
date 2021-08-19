package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jaskeerat789/controlers"
)

var port = os.Getenv("PORT")

func main() {
	l := hclog.Default()
	r := mux.NewRouter()
	handlers.CORS()

	pc := controlers.NewProductController()
	sc := controlers.NewStatusCotroller()
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: controlers.JwtMiddlewareHandler,
	})

	r.Handle("/status", jwtMiddleware.Handler(http.HandlerFunc(sc.GetStatus))).Methods("GET")
	r.Handle("/products", jwtMiddleware.Handler(http.HandlerFunc(pc.GetProduct))).Methods("GET")
	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(http.HandlerFunc(pc.GetProductBySlug))).Methods("POST")

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
