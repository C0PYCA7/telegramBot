package main

import (
	"goland/bot"
)

func main() {
	bot.StarBot()
	bot.StartDb()

	/*bot, err := tgbotapi.NewBotAPI("6219475155:AAGeqtKH--UY4suqHR9UtyTe8M7eufzS4Aw")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// send a reply message to the user
		replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Как тебя зовут?")
		_, err := bot.Send(replyMsg)
		if err != nil {
			log.Println(err)
		}

		// wait for the user's response
		name := ""
		for name == "" {
			update := <-updates
			if update.Message == nil {
				continue
			}
			name = update.Message.Text
		}

		// send a reply message to the user with the next question
		replyMsg = tgbotapi.NewMessage(update.Message.Chat.ID, "Сколько тебе лет?")
		_, err = bot.Send(replyMsg)
		if err != nil {
			log.Println(err)
		}

		// wait for the user's response
		age := ""
		for age == "" {
			update := <-updates
			if update.Message == nil {
				continue
			}
			age = update.Message.Text
		}

		// send a reply message to the user with their name and age
		replyMsg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+name+", тебе "+age+" лет.")
		_, err = bot.Send(replyMsg)
		if err != nil {
			log.Println(err)
		}
	}*/
}
