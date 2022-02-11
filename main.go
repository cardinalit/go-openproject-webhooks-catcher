package main

import (
   "net/http"

   bot "go-openproject-webhooks-bot/openproject-webhooks-bot"
)

func main() {
   http.HandleFunc("/openproject", bot.ServeHTTP)

   err := http.ListenAndServe(":80", nil)
   if err != nil {
      return
   }
}
