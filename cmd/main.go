package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/roman127001/urlsha/internal/shorter"
	store "github.com/roman127001/urlsha/internal/store/inruntime"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
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

	mux.HandleFunc("POST /encode", func(w http.ResponseWriter, r *http.Request) {
		type RS struct {
			ShortURL string `json:"short_url"`
			Error    string `json:"error"`
		}

		// Validate origin URL
		originURL := r.FormValue("url")

		type RQ struct {
			URL string `json:"url"`
		}
		var rq RQ
		if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
			log.Err(err)
		}
		originURL = rq.URL

		if len(originURL) == 0 {
			out, err := json.Marshal(RS{Error: "URL is empty"})
			if err != nil {
				log.Err(err)

			}
			if _, err := w.Write(out); err != nil {
				log.Err(err)
			}
			return
		}
		_, err := url.ParseRequestURI(originURL)
		if err != nil {
			log.Err(err)
		}

		// Assembling short URL
		key := sh.Generate()
		shortURL := schema + host + ":" + port + "/decode/" + key

		// Save short URL
		st.Set(key, originURL)

		// Prepare response
		out, err := json.Marshal(RS{ShortURL: shortURL})
		if err != nil {
			log.Err(err)
		}

		// Write response
		// TODO The output must be in json format?
		if _, err := w.Write(out); err != nil {
			log.Err(err)
		}
	})

	// TODO opened questions for decode:
	// 1. Just decode - without redirect (why); may be add another handler with redirect?
	// 2. Key must be determined by input data or random (random is more secure)?
	// ```
	// The service should generate short aliases for long URLs
	// and be capable of mapping those aliases back to the original URLs.
	// ```
	mux.HandleFunc("GET /decode/{short}", func(w http.ResponseWriter, r *http.Request) {
		type RS struct {
			OriginURL string `json:"origin_url"`
			Error     string `json:"error"`
		}

		shortURL := r.PathValue("short")

		// Get original URL from store
		originURL, ok := st.Get(shortURL)
		if !ok {
			out, err := json.Marshal(RS{OriginURL: "", Error: "URL not found"})
			if err != nil {
				log.Err(err)
			}
			if _, err := w.Write(out); err != nil {
				log.Err(err)
			}

			return
		}

		// Prepare response
		out, err := json.Marshal(RS{OriginURL: originURL})
		if err != nil {
			log.Err(err)
		}

		// Write response
		// TODO The output must be in json format?
		if _, err := w.Write(out); err != nil {
			log.Err(err)
		}
	})

	// `favicon.ico` to prevent 404 errors in the browser
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("")); err != nil {
			log.Err(err)
		}
	})

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
