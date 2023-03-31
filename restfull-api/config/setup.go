package config

import (
	"fmt"
	"github.com/fernanda-one/golang_api/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDatabase() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to Load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	//dbProv := os.Getenv("DB_PROV")
	//if dbProv == "mysql" {
	//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	//}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entities.User{}, &entities.Book{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPG, err := db.DB()
	if err != nil {
		panic("Failed to close database")
	}
	dbPG.Close()
}
