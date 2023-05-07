package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var bot, _ = tgbotapi.NewBotAPI("6219475155:AAGeqtKH--UY4suqHR9UtyTe8M7eufzS4Aw")
var step1, step2, step3 string

func StarBot() {
	_, err := tgbotapi.NewBotAPI("6219475155:AAGeqtKH--UY4suqHR9UtyTe8M7eufzS4Aw")

	if err != nil {
		log.Printf("Ошибка: %s", err.Error())
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // игнорировать любые обновления, которые не являются сообщениями
			continue
		}

		var tgID = update.Message.Chat.ID

		switch update.Message.Text {
		case "/start":
			StartHandler(update)
		case "/set":
			FindUser(tgID)
			SetHandler(update)
			step1 = ""
			for step1 == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				step1 = update.Message.Text
			}
			LoginHandler(update)
			step2 = ""
			for step2 == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				step2 = update.Message.Text
			}
			PassHandler(update)
			step3 = ""
			for step3 == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				step3 = update.Message.Text
			}
			SetRequest(step1, step2, step3, tgID)

		case "/get":
			GetHandler(update)
			step1 = ""
			for step1 == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				step1 = update.Message.Text
			}
			GetRequest(step1, tgID, update)
		case "/del":
			DelHandler(update)
			step1 = ""
			for step1 == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				step1 = update.Message.Text
			}
			DelRequest(step1, tgID)
		}
	}
}
