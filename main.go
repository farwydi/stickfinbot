package stickfinbot

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
}

func TelegramHook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Fail read body from request: %v", err)
		http.Error(w, "Bad body data", http.StatusBadRequest)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Printf("Fail unmarshal json: %v", err)
		http.Error(w, "Fail read json", http.StatusBadRequest)
		return
	}

	// ignore any non-Message Updates
	if update.Message == nil {
		http.Error(w, "Not support method", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("Fail send message: %v", err)
		http.Error(w, "Bot fail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, "OK")
	if err != nil {
		log.Printf("Fail send ok respons: %v", err)
	}
}
