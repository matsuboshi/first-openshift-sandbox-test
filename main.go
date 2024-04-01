package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprintf(w, "hello gopher")
		return
	}
	fmt.Fprintf(w, "404 Not Found")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
