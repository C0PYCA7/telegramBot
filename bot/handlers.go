package bot

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func StartHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, ""+
		"Привет\nДанный бот предназначен для хранения данных "+
		"от социальных сетей в нем доступен следующий список команд:\n"+
		"/set - добаляет логин и пароль к сервису\n"+
		"/get - получает логин и пароль по названию сервиса\n"+
		"/del - удаляет значения для сервиса")
	bot.Send(msg)
}

func SetHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Какому сервису вы хотите добавить данные?")
	bot.Send(msg)
}

func LoginHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ваш логин: ")
	bot.Send(msg)
}

func PassHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ваш пароль: ")
	bot.Send(msg)
}

func GetHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите название сервиса: ")
	bot.Send(msg)
}

func DelHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите название сервиса: ")
	bot.Send(msg)
}

func DeletePasswordAfterDelay(bot *tgbotapi.BotAPI, chatID int64, messageID int, delay time.Duration) {
	time.Sleep(delay)

	deleteMsg := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}

	_, err := bot.DeleteMessage(deleteMsg)
	if err != nil {
		fmt.Println("Ошибка удаления сообщения:", err)
	}
}
