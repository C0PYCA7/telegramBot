package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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

/*func FindService(step1 string) {
	query := "SELECT COUNT(*) FROM services WHERE name = ?"
	var count int
	err := db.QueryRow(query, step1).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$", count)

	if count == 0 {
		insertQuery := "INSERT INTO services (name) VALUES (?)"
		_, err = db.Exec(insertQuery, step1)
		if err != nil {
			log.Fatal(err)
		}
	}
}*/

func SetRequest(step1 string, step2 string, step3 string, tg_id int64) error {

	var userId int

	uId := "SELECT id FROM users WHERE tg_id = ?"
	err = db.QueryRow(uId, tg_id).Scan(&userId)
	if err != nil {
		log.Fatal(err)
	}

	sQuery := "SELECT COUNT(*) FROM services WHERE name = ?"
	var count int
	err := db.QueryRow(sQuery, step1).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		insertQuery := "INSERT INTO services (name) VALUES (?)"
		_, err = db.Exec(insertQuery, step1)
		if err != nil {
			log.Fatal(err)
		}
	}

	query := "SELECT * FROM user_info " +
		"LEFT JOIN users ON user_info.user_id = users.id " +
		"LEFT JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"
	rows, err := db.Query(query, tg_id, step1)

	if !rows.Next() {
		query = "INSERT INTO user_info (user_id, service_id, login, password)" +
			" VALUES ((SELECT id FROM users WHERE tg_id = ?)," +
			"(SELECT id FROM services WHERE name = ?),?,?)"
		_, err = db.Exec(query, tg_id, step1, step2, step3)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = "UPDATE user_info " +
			"JOIN users ON user_info.user_id = users.id " +
			"JOIN services ON user_info.service_id = services.id " +
			"SET user_info.user_id = users.id, user_info.service_id = services.id, login = ?, password = ? " +
			"WHERE users.tg_id = ? AND services.name = ?"
		_, err = db.Exec(query, step2, step3, tg_id, step1)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func GetRequest(step1 string, tg_id int64, update tgbotapi.Update) {
	query := "SELECT login, password FROM user_info " +
		"JOIN users ON user_info.user_id = users.id " +
		"JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"

	rows, err := db.Query(query, tg_id, step1)
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
		_, err = bot.Send(msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func DelRequest(step1 string, tg_id int64) {
	query := "DELETE user_info FROM user_info " +
		"JOIN users ON user_info.user_id = users.id " +
		"JOIN services ON user_info.service_id = services.id " +
		"WHERE users.tg_id = ? AND services.name = ?"

	_, err = db.Exec(query, tg_id, step1)
	if err != nil {
		log.Fatal(err)
	}
}
