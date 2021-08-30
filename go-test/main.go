package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hii, this is v1.0.0")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":6111", nil)
}
