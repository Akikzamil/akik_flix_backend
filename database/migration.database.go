package database

import (
	"akikflix/model"
	"log"
)

func migrateDatabase() {
	log.Println("Migrating database...")
	DB.AutoMigrate(&model.User{},&model.Video{},&model.Playlist{})
}
