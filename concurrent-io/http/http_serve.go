package http

import (
	"fmt"
	h "net/http"
)

const LISTEN_ADDR = ":8080"

var Content string = "Hello world!"

func hello(w h.ResponseWriter, req *h.Request) {
	fmt.Fprintf(w, "%s", Content)
}

func PrintServerInfo() {
	fmt.Printf("HTTP Server Listening on %s\n", LISTEN_ADDR)
}

func ServeHttp() {
	h.HandleFunc("/", hello)
	h.ListenAndServe(LISTEN_ADDR, nil)
}
