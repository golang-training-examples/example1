package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Hostname string `json:"hostname"`
}

func Server(port int) {
	HOSTNAME, _ := os.Hostname()

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Hostname: HOSTNAME})
	})

	fmt.Printf("Listening on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
