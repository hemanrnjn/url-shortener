package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hemanrnjn/url-shortner/app"
	u "github.com/hemanrnjn/url-shortner/utils"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)
	originalUrl := app.OriginalURL(url["url"])
	if originalUrl == "" {
		http.Redirect(w, r, "Not Found", 404)
	} else {
		http.Redirect(w, r, app.OriginalURL(url["url"]), 301)
	}
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {

	var url app.URL

	err := json.NewDecoder(r.Body).Decode(&url) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	shortURL := app.ShortenedURL(url.URL)
	log.Println("Shortened URL: ", shortURL)

	response := u.Message(true, "URL Shortened")
	response["url"] = shortURL
	u.Respond(w, response)
}
