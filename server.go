package main

import (
	"fmt"
	"net/http"
	"os"
)

func startDummyServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Helsa bridge running")
	})

	go http.ListenAndServe(":"+port, nil)
}
