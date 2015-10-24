package main

import (
  "net/http"
  "log"
  "encoding/json"
)

// The Thing Model, mapping the URL.Query() contents
type Thing struct {
  Values map[string][]string
}

// Handles get requests for the /things service
func things(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  thing := Thing{Values: query}
  json.NewEncoder(w).Encode(thing)
}

func main() {
  http.HandleFunc("/things", things)
  log.Fatal(http.ListenAndServe(":9090", nil))
}
