package main

import (
	"fmt"
	"github.com/ZaharBorisenko/tg-bot/clients/openweather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	Client := openweather.New(os.Getenv("secret_key_weather"))

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			coord, err := Client.GetCoordinates(update.Message.Text)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, вы точно указали верный город?")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Долгота %f, Широта %f", coord.Lon, coord.Lat))
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
