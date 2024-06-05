package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Vercel!")
}

func main() {
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":3000", nil)
}
