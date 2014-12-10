package main

import (
    // "log"
    "net/http"
    "fmt"
    // "io"
    // "encoding/json"
)

func main() {
    fmt.Printf("Hello")
    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Result struct {
    Msg string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "application/json")
    // result, _ := json.Marshal(Result{"hello Jobr"})
    // io.WriteString(w, string(result))
    fmt.Fprint(w, "Hello, world!")
}