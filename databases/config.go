package databases

import (
	"Golang10/Final/Ardi/models"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Username: "root",   //database username
			Password: "",       //database password
			Name:     "golang", //database name
		},
	}
}

func SetupDatabase(config *Config) {
	configDb := os.Getenv("DATABASE_URL")
	configDb = strings.ReplaceAll(configDb, "postgres://", "")
	user := strings.Split(configDb, ":")
	pw := strings.Split(user[1], "@")
	port := strings.Split(user[2], "/")
	msgArgs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", pw[1], port[0], user[0], pw[0], port[1], "disable", "5")
	db, err := gorm.Open(config.DB.Dialect, msgArgs)
	fmt.Println(msgArgs)

	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect to mysql")
		return
	}

	fmt.Println("Connected to mysql")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.About{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Contact{})
}
