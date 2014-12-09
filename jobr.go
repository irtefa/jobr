package main

import (
    "log"
    "net/http"
    "io"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", HandleJSON)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Result struct {
    Msg string `json:"msg"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(Result{"hello Jobr"})
    io.WriteString(w, string(result))
}