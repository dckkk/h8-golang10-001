package databases

import (
	"Golang10/Final/Ardi/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
			Dialect:  "mysql",
			Host:     "localhost",
			Username: "root",   //database username
			Password: "",       //database password
			Name:     "golang", //database name
		},
	}
}

func SetupDatabase(config *Config) {
	db, err := gorm.Open(config.DB.Dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

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
