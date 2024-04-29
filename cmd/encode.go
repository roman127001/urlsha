package main

import (
	"encoding/json"
	"github.com/roman127001/urlsha/internal/shorter"
	store "github.com/roman127001/urlsha/internal/store/inruntime"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
)

func encodeHandler(sh shorter.Shorter, st store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
