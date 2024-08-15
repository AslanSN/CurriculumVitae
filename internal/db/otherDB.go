package db

// import (
// 	"log"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// /**
// // For the main.go:
// Imports

// "github.com/AslanSN/CurriculumVitae/internal/db"
// "github.com/AslanSN/CurriculumVitae/internal/db/models"

// db.DBconnection()

// db.DB.AutoMigrate(models.AboutMeType{}, models.ExperienceStruct{})
// */

// var DB *gorm.DB

// func DBconnection() {
// 	ip := "172.17.0.2"
// 	dsn := "host=" + ip + " user=aslan password=ThisIsMyEspAÑisPassword dbname=cv port=5432"

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	} else {
// 		log.Println("DB connected", DB)

// 	}

// }
