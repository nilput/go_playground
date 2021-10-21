package main

import (
	"concurrent_io/cmd"
	"concurrent_io/http"
)

func main() {
	stdinControlChan := make(chan bool)
	http.PrintServerInfo()
	go cmd.StdinControl(stdinControlChan)
	http.ServeHttp()
}
