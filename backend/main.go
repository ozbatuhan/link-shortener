
package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "encoding/json"
)

type User struct {
    Username string
    Password string
}

type UrlStruct struct {
    LongURL  string
    ShortURL string
}

var users = make(map[string]User)
var store = make(map[string]UrlStruct)
var hasher = bcrypt.DefaultCost

func main() {
    http.HandleFunc("/", UrlRedirect)
    http.HandleFunc("/create", AuthenticatedShortenUrl)
    log.Println("Listening on port 8090")
    http.ListenAndServe(":8090", nil)
}

func UserAuthenticated(r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        return false
    }

    storedUser, exists := users[username]
    if !exists {
        return false
    }

    err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))
    return err == nil
}

func AuthenticatedShortenUrl(w http.ResponseWriter, r *http.Request) {
    if !UserAuthenticated(r) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

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

func UrlRedirect(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[1:]
    if urlStruct, ok := store[shortURL]; ok {
        http.Redirect(w, r, urlStruct.LongURL, http.StatusSeeOther)
    } else {
        http.NotFound(w, r)
    }
}

func generateShortURL() string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(length(letters))]
    }
    return string(b)
}
