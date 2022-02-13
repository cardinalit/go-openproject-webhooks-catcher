package main

import (
   "net/http"

   bot "go-openproject-webhooks-bot/openproject-webhooks-bot"
)

func main() {
   http.HandleFunc("/wp/created", bot.ServeHTTP)
   // TODO: Сделать нотификацию на обновление задачи
   // http.HandleFunc("/wp/updated", bot.ServeHTTP)

   err := http.ListenAndServe(":80", nil)
   if err != nil {
      return
   }
}
