package database

import (
	"fmt"
	"os"

	"github.com/AuthSystemJWT/deyki/v2/util"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type DBConfig struct {
	Host string
	User string
	Password string
	Name string
	Port string
}


type User struct {
	gorm.Model
	Username	string	`gorm:"unique" json:"username"`
	Password	string	`json:"password"`
}


func LoadEnvVariables() *util.ErrorMessage {

	errorMessage := godotenv.Load(".env")
	if errorMessage != nil {
		return util.ErrorMessage{}.ErrorLoadingEnvFile()
	}

	return nil
}


func ConnectDB() (*gorm.DB, *util.ErrorMessage) {

	dbConfig := &DBConfig{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	db.AutoMigrate(&User{})

	return db, nil
}
