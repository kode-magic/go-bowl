package demographics

import (
	"database/sql"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDemograpgicsDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./demographics/demographics.db")
	CheckErr(err)

	return db
}

func ConnectGormDB() (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open("./demographics/demographics.db"), &gorm.Config{})
	CheckErr(err)

	db.AutoMigrate(
		&Country{},
		&Region{},
		&District{},
	)
  return db, nil
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}