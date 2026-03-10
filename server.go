package main

import (
	"net/http"
	"os"
)

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helsa bridge running"))
	})

	go http.ListenAndServe(":"+port, nil)
}
