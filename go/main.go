package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-APP-USER-ID", "12345")
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}