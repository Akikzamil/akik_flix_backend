package database

import (
	"akikflix/util"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := getDsn();
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database")

	migrateDatabase();
}

func getDsn() string {
	dbHost := util.GetVariable("dbHost")
	dbUser := util.GetVariable("dbUser")
	dbPassword := util.GetVariable("dbPassword")
	dbDbname := util.GetVariable("dbDbname")
	dbPort := util.GetVariable("dbPort")
	dbSslmode := util.GetVariable("dbSslmode")
	
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",dbHost,dbUser,dbPassword,dbDbname,dbPort,dbSslmode)
}
