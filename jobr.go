package main

import (
    "net/http"
    "os"
    "io"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", handler)
    
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

type Result struct {
    Msg string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(Result{"hello Jobr"})
    io.WriteString(w, string(result))
}
