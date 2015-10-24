package main

import (
  "net/http"
  "log"
  "encoding/json"
)

type Thing struct {
  Name  string
}

func things(w http.ResponseWriter, r *http.Request) {
  thing := Thing{Name: "A thing"}
  json.NewEncoder(w).Encode(thing)
}

func main() {
  http.HandleFunc("/things", things)
  log.Fatal(http.ListenAndServe(":9090", nil))
}
