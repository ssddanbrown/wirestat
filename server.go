package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type middlewareFunc func(http.Handler) http.Handler

func startServer(respBuilder ResponseFunc, port uint, accessKey string) {
	handler := getServerHandler(respBuilder, accessKey)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler)
}

func getServerHandler(respBuilder ResponseFunc, accessKey string) http.Handler {
	mux := http.NewServeMux()

	accessMiddleware := getAccessControlMiddleware(accessKey)
	responseHandler := getResponseHandler(respBuilder)

	mux.Handle("/", accessMiddleware(responseHandler))

	return mux
}

func getAccessControlMiddleware(accessKey string) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if accessKey != "" {
				keyHeader := r.Header.Get("X-Access-Key")
				keyQuery := r.URL.Query().Get("key")
				if accessKey != keyHeader && accessKey != keyQuery {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid access key provided"))
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getResponseHandler(builder ResponseFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := builder()

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
	})
}
