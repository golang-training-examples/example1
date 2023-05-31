package server

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Response struct {
	Hostname string `json:"hostname"`
}

func Server(port int) {
	HOSTNAME, _ := os.Hostname()

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", http.StatusNotFound).
			Send()
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", http.StatusOK).
			Send()
		json.NewEncoder(w).Encode(Response{Hostname: HOSTNAME})
	})

	log.Info().
		Msgf("Listening on 0.0.0.0:%d, see http://127.0.0.1:%d", port, port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
