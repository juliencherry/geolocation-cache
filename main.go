package main

import (
   "fmt"
   "log"
   "net/http"
)

import "github.com/juliencherry/geolocation-cache/api"

func main() {
   http.HandleFunc("/", api.Handler);

   fmt.Println("Starting server at port 8080")
   if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal(err)
   }
}
