package database

import (
	"log"
	"markdown-note/internal/models/note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&note.Note{})
	return db
}
