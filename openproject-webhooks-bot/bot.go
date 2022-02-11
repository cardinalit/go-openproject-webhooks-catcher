package openproject_webhooks_bot

import (
   "fmt"
   "net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")

   _, err := fmt.Fprintf(w, "OK\n")
   if err != nil {
      return
   }
}
