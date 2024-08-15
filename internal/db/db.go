package db

import (
	"log"
	"os"

	models "github.com/AslanSN/CurriculumVitae/internal/db/models"
	storage "github.com/AslanSN/CurriculumVitae/internal/db/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnection() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return
	}

	// err = models.MigrateExperience(db)
	err = db.AutoMigrate(&models.ExperienceStruct{}, &models.AboutMeType{})

	if err != nil {
		log.Fatal("Could not migrate the database. ", err)
		return
	}

	DB = db

	log.Println("DB connected successfully 🤙")
}
