package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
