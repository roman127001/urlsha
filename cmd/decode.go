package main

import (
	"encoding/json"
	store "github.com/roman127001/urlsha/internal/store/inruntime"
	"github.com/rs/zerolog/log"
	"net/http"
)

func decodeHandler(st store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
