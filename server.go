package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func startServer() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	system := GetLatestSystem()

	rJson, err := json.MarshalIndent(system, "", "    ")
	if err != nil {
		rJson = []byte(fmt.Sprintf("Failed to properly encode system data to JSON, with error: %s\n", err))
	}

	w.Header().Set("Content-Type", "application/json")

	if len(system.Alerts) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(rJson)
}
