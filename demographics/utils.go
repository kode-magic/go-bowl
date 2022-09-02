package demographics

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectGormDB() (*gorm.DB, error) {
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
