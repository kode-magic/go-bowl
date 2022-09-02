package demographics

import (
	"database/sql"
	"log"
)

func connectDemograpgicsDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./demographics/demographics.db")
	CheckErr(err)
	// defer close
	// defer db.Close()

	return db
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}