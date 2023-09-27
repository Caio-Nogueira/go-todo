package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"github.com/Caio-Nogueira/go-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Hostname string `json:"hostname"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func Connect() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic("Could not open config file")
	}
	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic("Could not read config file")
	}

	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		panic("Could not parse config file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Hostname, config.Username, config.Password, config.Database, config.Port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	
	db.Debug().AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})

	DB = db

}


