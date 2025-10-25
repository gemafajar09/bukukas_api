package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      int
	DBName      string
	JWT         string
	GIN_MODE    string
	TIME_FORMAT string
	TIME_ZONE   string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env vars")
	}

	viper.SetDefault("SERVER_PORT", ":8080")

	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_PORT", 3306)
	viper.SetDefault("DB_NAME", "mydb")

	viper.SetDefault("JWT_SECRET", "secret_key")
	viper.SetDefault("GIN_MODE", "")

	viper.SetDefault("TIME_FORMAT", "02-01-2006 15:04:05")
	viper.SetDefault("TIME_ZONE", "Asia/Jakarta")

	viper.AutomaticEnv()

	return Config{
		ServerPort:  viper.GetString("SERVER_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetInt("DB_PORT"),
		DBName:      viper.GetString("DB_NAME"),
		JWT:         viper.GetString("JWT_SECRET"),
		GIN_MODE:    viper.GetString("GIN_MODE"),
		TIME_FORMAT: viper.GetString("TIME_FORMAT"),
		TIME_ZONE:   viper.GetString("TIME_ZONE"),
	}
}
