package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {

	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	createTable := `CREATE TABLE IF NOT EXISTS notes (
	"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"word" TEXT,
	"defination" TEXT,
	"category" TEXT
	);`

	statement, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("notes table created successfully")
}

func InsertNote(word string, defination string, category string) {
	insertNoteSql := `INSERT INTO notes(word , defination , category)
	VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertNoteSql)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(word, defination, category)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("inserted into notes successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM notes ORDER BY word")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "—", definition)
	}
}
