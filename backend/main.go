
package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "net/http/httputil"
    "encoding/json"
)

type UrlStruct struct {
    LongURL  string
    ShortURL string
}

var store = make(map[string]UrlStruct)

func main() {
    http.HandleFunc("/", UrlRedirect)
    http.HandleFunc("/create", ShortenUrl)
    log.Println("Listening on port 8090")
    http.ListenAndServe(":8090", nil)
}

func UrlRedirect(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[1:]
    if urlStruct, ok := store[shortURL]; ok {
        http.Redirect(w, r, urlStruct.LongURL, http.StatusSeeOther)
    } else {
        http.NotFound(w, r)
    }
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
    var url UrlStruct
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    shortURL := generateShortURL()
    url.ShortURL = shortURL
    store[shortURL] = url
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(url)
}

func generateShortURL() string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(lengletters))]
    }
    return string(b)
}
