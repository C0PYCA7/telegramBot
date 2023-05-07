package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func FindUser(tg_id int64) {

	query := "SELECT tg_id FROM users WHERE tg_id = ?"
	rows, err := db.Query(query, tg_id)
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		insertQuery := "INSERT INTO users (tg_id) VALUES (?)"
		_, err = db.Exec(insertQuery, tg_id)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func SetRequest(servName string, uLogin string, uPass string, tg_id int64) error {

	var userId int

	uId := "SELECT id FROM users WHERE tg_id = ?"
	err := db.QueryRow(uId, tg_id).Scan(&userId)
	if err != nil {
		log.Fatal(err)
	}

	sQuery := "SELECT COUNT(*) FROM services WHERE name = ?"
	var count int
	err = db.QueryRow(sQuery, servName).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		insertQuery := "INSERT INTO services (name) VALUES (?)"
		_, err = db.Exec(insertQuery, servName)
		if err != nil {
			log.Fatal(err)
		}
	}

	query := "SELECT * FROM user_info " +
		"LEFT JOIN users ON user_info.user_id = users.id " +
		"LEFT JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"
	rows, err := db.Query(query, tg_id, servName)

	if !rows.Next() {
		query = "INSERT INTO user_info (user_id, service_id, login, password)" +
			" VALUES ((SELECT id FROM users WHERE tg_id = ?)," +
			"(SELECT id FROM services WHERE name = ?),?,?)"
		_, err = db.Exec(query, tg_id, servName, uLogin, uPass)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = "UPDATE user_info " +
			"JOIN users ON user_info.user_id = users.id " +
			"JOIN services ON user_info.service_id = services.id " +
			"SET user_info.user_id = users.id, user_info.service_id = services.id, login = ?, password = ? " +
			"WHERE users.tg_id = ? AND services.name = ?"
		_, err = db.Exec(query, uLogin, uPass, tg_id, servName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func GetRequest(servName string, tg_id int64, update tgbotapi.Update) {
	query := "SELECT login, password FROM user_info " +
		"JOIN users ON user_info.user_id = users.id " +
		"JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"

	rows, err := db.Query(query, tg_id, servName)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var login, password string
		err := rows.Scan(&login, &password)
		if err != nil {
			log.Fatal(err)
		}

		userInfo := login + " " + password

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, userInfo)
		sentMsg, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		delay := 1 * time.Minute
		go DeletePasswordAfterDelay(bot, update.Message.Chat.ID, sentMsg.MessageID, delay)
	}
}

func DelRequest(servName string, tg_id int64) {
	query := "DELETE user_info FROM user_info " +
		"JOIN users ON user_info.user_id = users.id " +
		"JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"

	_, err := db.Exec(query, tg_id, servName)
	if err != nil {
		log.Fatal(err)
	}
}
