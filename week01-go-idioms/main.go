package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/long", LongPooling)
	if err := http.ListenAndServe("8080", nil); err != nil {
		panic(err)
	}

}
