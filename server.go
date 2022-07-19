package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var responseBuilder *ResponseBuilder

func startServer(respBuilder *ResponseBuilder) {
	responseBuilder = respBuilder

	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	response := responseBuilder.GetResponseData()

	rJson, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		rJson = []byte(fmt.Sprintf("Failed to properly encode system data to JSON, with error: %s\n", err))
	}

	w.Header().Set("Content-Type", "application/json")

	if len(response.Alerts) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(rJson)
}
