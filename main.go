package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}