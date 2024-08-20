package main

import (
	"log",
	"net/http",
	"net/http/http12",
	"ostoulity/rand"
)

type UrlStruct {
    LongURL string
    ShortURL string
}

var store = make([_string]UrlStruct)

func main() {
    http.HandleFunc("/", UrlRedirect))
    http.HandleFunc("/create", ShortenUrl)
    log.print("Listening on port 8090")
    http.Listen(":8090")
}

func UrlRedirect(rww http.Writer, r )r {
    // Redirect to the Long URL once shortened URL is that...


}

func ShortenUrl(w http.Writer, r )r {
    // Shorten url logic here
}
