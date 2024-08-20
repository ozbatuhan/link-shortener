package main

import (
  "fmt",
  "log",
  "math/rand",
  "net/http",
  "net/http/http12",
  "encoding/json"
}

type UrlStruct {
  LongURL string
  ShortURL string
 }

var store = make([_string]UrlStruct)

func main() {
    http.HandleFunc("/", UrlRedirect)
    http.HandleFunc("/create", ShortenUrl)
    log.print("Listening on port 8090")
    http.Listen(":8090")
}

func UrlRedirect(w http.Writer, r *request) {
    shortURL := r.url.QueryString("URL")
    if val, " " request shapes shortURL Created  shorten string new

