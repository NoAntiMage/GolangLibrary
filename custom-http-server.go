package main

import (
	"fmt"
	"net/http"
	"time"
)

type textHandler struct {
	responseText string
}

// implement ServeHTTP of Handler
func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	// self-defined Multiplexer
	m := http.NewServeMux()

	m.HandleFunc("/ping", Ping)

	t := &textHandler{"interface Handler"}
	m.Handle("/handler", t)

	// self-defined HTTPserver
	server := &http.Server{
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	server.Handler = m
	server.ListenAndServe()
}