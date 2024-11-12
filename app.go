package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola Mundo en Go!")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8001", nil)
}