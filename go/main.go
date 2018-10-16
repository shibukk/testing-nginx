package main

import (
  "fmt"
  "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-APP-USER-ID", "12345")
  fmt.Fprintf(w, "Home")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Login")
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/login", loginHandler)
  http.ListenAndServe(":8080", nil)
}