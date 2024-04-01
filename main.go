package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "hello earthlings")
		return
	}
	fmt.Fprintf(w, "404 Not Found")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
