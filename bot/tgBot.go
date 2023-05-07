package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
	"time"
)

var bot, _ = tgbotapi.NewBotAPI("6219475155:AAGeqtKH--UY4suqHR9UtyTe8M7eufzS4Aw")

type Session struct {
	ServName string
	ULogin   string
	UPass    string
	Mutex    sync.Mutex
}

type User struct {
	ChatID        int64
	PasswordMsgId int
	Session       *Session
}

func StarBot() {
	servName, uLogin, uPass := "", "", ""
	_, err := tgbotapi.NewBotAPI("6219475155:AAGeqtKH--UY4suqHR9UtyTe8M7eufzS4Aw")

	if err != nil {
		log.Printf("Ошибка: %s", err.Error())
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	users := make(map[int64]*User)

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		tgID := update.Message.Chat.ID
		chatID := update.Message.Chat.ID

		user, ok := users[chatID]
		if !ok {
			user = &User{
				ChatID: chatID,
			}
			users[tgID] = user
		}

		switch update.Message.Text {
		case "/start":
			StartHandler(update)
		case "/set":

			FindUser(tgID)
			SetHandler(update)

			user.Session = &Session{}

			servName = ""
			for servName == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				servName = update.Message.Text
			}
			user.Session.ServName = servName

			LoginHandler(update)
			uLogin = ""
			for uLogin == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				uLogin = update.Message.Text
			}
			user.Session.ULogin = uLogin

			PassHandler(update)
			uPass = ""
			for uPass == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				uPass = update.Message.Text
				user.PasswordMsgId = update.Message.MessageID
			}
			user.Session.UPass = uPass
			SetRequest(user.Session.ServName, user.Session.ULogin, user.Session.UPass, tgID)
			delay := 1 * time.Minute
			go DeletePasswordAfterDelay(bot, tgID, user.PasswordMsgId, delay)

		case "/get":
			GetHandler(update)
			servName = ""
			for servName == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				servName = update.Message.Text
			}
			GetRequest(servName, tgID, update)

		case "/del":
			DelHandler(update)
			servName = ""
			for servName == "" {
				update := <-updates
				if update.Message == nil {
					continue
				}
				servName = update.Message.Text
			}
			DelRequest(servName, tgID)
		}
	}
}
