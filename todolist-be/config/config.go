package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port     string
	Driver   string
	Name     string
	Address  string
	DB_Port  string
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("can't read file env: ", err.Error())
	}

	var defaultConfig AppConfig
	defaultConfig.Port = os.Getenv("PORT")
	defaultConfig.Driver = os.Getenv("DRIVER")
	defaultConfig.Name = os.Getenv("DB_NAME")
	defaultConfig.Address = os.Getenv("ADDRESS")
	defaultConfig.DB_Port = os.Getenv("DB_PORT")
	defaultConfig.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Password = os.Getenv("DB_PASSWORD")

	return &defaultConfig
}
