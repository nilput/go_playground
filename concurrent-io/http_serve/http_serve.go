package http_serve

import (
	"fmt"
	"net/http"
)

var content string = "Hello world!"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s", content)
}

func ServeHttp() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
