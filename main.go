package main

import (

    "fmt"
    "log"
    "net/http"
)


func main() {
    http.HandleFunc("/", handlerFunc)
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    log.Printf("Ping from %s", r.RemoteAddr)
    fmt.Fprintln(w, "Hello World! from OpenShift")
}