package bot

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db, err = sql.Open("mysql", "Alvarez:89173036695Druno@tcp(127.0.0.1:3306)/logPassManager")

func StartDb() {
	db, err := sql.Open("mysql", "Alvarez:89173036695Druno@tcp(127.0.0.1:3306)/logPassManager")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS logpassmanager")
	if err != nil {
		log.Fatal(err)
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
