package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	h "github.com/hemanrnjn/url-shortner/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{url}", h.RedirectHandler)
	r.HandleFunc("/shorten/", h.ShortenHandler).Methods("POST")

	log.Println("Server Running and listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
