package main

import "net/http"
import "webhook-alertmanager-logger/api/v1"

func main() {
	http.HandleFunc("/api/v1/logger", func(w http.ResponseWriter, r *http.Request) {
		api.Logger(w, r)
	})

	http.ListenAndServe(":9092", nil)
}
