package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to run sever", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
}
