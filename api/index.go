package api

import (
    "io/ioutil"
    "log"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

   apiKeyValues, ok := r.URL.Query()["apiKey"]
   if !ok || len(apiKeyValues) != 1 {
      log.Println("Unable to parse query")
      return
   }

   ipValues, ok := r.URL.Query()["ip"]
   if !ok || len(ipValues) != 1 {
      log.Println("Unable to parse query")
      return
   }

   resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=" + apiKeyValues[0] + "&ip=" + ipValues[0])
   if err != nil {
      log.Println("Unable to get response from geolocation API")
      return
   }

   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Println("Unable to read response body from geolocation API")
      return
   }

   w.Header().Set("Content-Type", "application/json")
   _, err = w.Write(body)
   if err != nil {
      log.Println("Unable to write response")
   }
}
