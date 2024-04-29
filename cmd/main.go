package main

import (
	"context"
	"errors"
	"github.com/roman127001/urlsha/internal/shorter"
	store "github.com/roman127001/urlsha/internal/store/inruntime"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO add authentication ?
// TODO add logging middleware ?
// TODO add error handling middleware ?
// TODO add expiration URLs?
// TODO add `author` for URL?
// TODO add `description` for URL?
// TODO use redis (or other outer storage)?
// TODO user Name or ID as prefix for short URL?

// TODO opened questions for decode:
// 1. Just decode - without redirect (why?). May be add another handler with redirect?
//
// ```
// The service should generate short aliases for long URLs
// and be capable of mapping those aliases back to the original URLs.
// ```
// 2. Key must be determined by input data or random? Random is more secure?

const (
	// TODO move those params to config!
	host   = "127.0.0.1"
	schema = "http://" // TODO use https
	port   = "8080"

	ctxTimeout = 5 * time.Second
)

func main() {
	sh := shorter.New()
	st := store.New()

	mux := http.NewServeMux()

	// TODO add v1 prefix
	mux.HandleFunc("POST /encode", encodeHandler(sh, st))
	mux.HandleFunc("GET /decode/{short}", decodeHandler(st))

	// `favicon.ico` to prevent 404 errors in the browser
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("")); err != nil {
			log.Err(err)
		}
	})

	// TODO add "/health", "version", "metrics" endpoints
	// TODO add http.Handle("/metrics", promhttp.Handler())

	// Run server
	srv := &http.Server{Addr: ":" + port, Handler: mux}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":"+port, mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err)
		}
	}()

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
