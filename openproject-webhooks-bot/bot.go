package openproject_webhooks_bot

import (
   "bytes"
   "encoding/json"
   "fmt"
   "html/template"
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "strconv"

   "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")

   _, err := fmt.Fprintf(w, "OK\n")
   if err != nil {
      return
   }

   b, _ := ioutil.ReadAll(r.Body)
   log.Println(string(b))

   var g workPackage

   err = json.Unmarshal(b, &g)
   if err != nil {
      log.Println(err)
   }

   switch r.URL.Path {
   case "/wp/created":
      chatID, _ := strconv.ParseInt(os.Getenv("OP_WEBHOOKS_BOT_TG_MAIN_CHAT_ID"), 10, 64)
      t := wp("created", &g)
      bot, _ := tgbotapi.NewBotAPI(os.Getenv("OP_WEBHOOKS_BOT_TG_BOT_KEY"))
      msg := tgbotapi.NewMessage(chatID, t)

      msg.ParseMode = "HTML"
      _, err = bot.Send(msg)
      if err != nil {
         log.Println(err)
      }
   }
}

func wp(action string, w *workPackage) string {
   var b bytes.Buffer

   tmpl, _ := template.ParseFiles("openproject-webhooks-bot/templates/message_work_package.html")
   if err := tmpl.Execute(&b, w); err != nil {
      fmt.Println(err)
   }

   return b.String()
}
