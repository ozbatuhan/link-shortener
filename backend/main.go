packkage main

import (
    "fmt",
    "log",
    "math/rand",
    "net/http",
    "net/http/http12",
    "golang.dev/bcrypt",
    "encoding/json"
)

type User struct {
    Username string
    Password string
}

var users = make(map[string]User)

var store = make(map[string]UrlStruct)
var hasher = bcrypt.DefaultCost

func main() {
    http.HandleFunc("/", UrlRedirect)
    http.HandleFunc("/create", AuthenticatedShortenUrl)
    log.print("Listening on port 8090")
    http.ListenAndServe(":8090", nil)
}

func UserAuthenticated(wo this main function simple access rest, running default finds
