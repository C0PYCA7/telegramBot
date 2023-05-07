package bot

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func StartDb() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(db:3306)/")

	err = db.Ping()
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS logpassmanager")
	if err != nil {
		log.Print("Wait a second")
	}

	db, err = sql.Open("mysql", "root:root@tcp(db:3306)/logpassmanager")
	err = db.Ping()
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("USE logpassmanager")
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users " +
		"(id INT AUTO_INCREMENT PRIMARY KEY," +
		"tg_id BIGINT UNIQUE);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS services(" +
		"id INT AUTO_INCREMENT PRIMARY KEY," +
		"name VARCHAR(100) UNIQUE);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user_info(" +
		"id INT AUTO_INCREMENT PRIMARY KEY," +
		"user_id INT," +
		"service_id INT," +
		"login VARCHAR(50)," +
		"password VARCHAR(50)," +
		"FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE," +
		"FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE);")
	if err != nil {
		log.Fatal(err)
	}
}
