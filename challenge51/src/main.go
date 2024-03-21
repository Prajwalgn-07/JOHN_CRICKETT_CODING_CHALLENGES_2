package main

import (
	"fmt"
	"net/http"
	"src/handlers"
)

func main() {
	proxyPort := "8989"
	http.HandleFunc("/", handlers.HandleRequestAndRedirect)
	http.ListenAndServe(":"+proxyPort, nil)
	fmt.Printf("Proxy server listening on port %s...\n", proxyPort)
}
