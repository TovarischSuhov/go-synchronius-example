package main

import (
	"net/http"

	"github.com/TovarischSuhov/go-synchronius-example/internal/server"
	"github.com/TovarischSuhov/log"
)

func main() {
	http.HandlerFunc("/ping", server.PingHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Error(err)
	}
}
