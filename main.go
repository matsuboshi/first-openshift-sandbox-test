package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello gopher 2")
	fmt.Println("hello from the terminal 2")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
