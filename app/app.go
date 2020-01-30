package app

import (
	"crypto/sha1"
	"encoding/base64"
)

var urlMap map[string]string

type URL struct {
	URL string `json:"url"`
}

func ShortenedURL(url string) string {
	urlMap = make(map[string]string)
	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	urlMap[url] = sha
	return sha
}

func OriginalURL(url string) string {
	for k, v := range urlMap {
		if v == url {
			return k
		}
	}
	return ""
}
