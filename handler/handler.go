package handler

import (
	"Golang10/Final/Ardi/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	configDb := os.Getenv("DATABASE_URL")
	configDb = strings.ReplaceAll(configDb, "postgres://", "")
	user := strings.Split(configDb, ":")
	pw := strings.Split(user[1], "@")
	port := strings.Split(user[2], "/")
	msgArgs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", pw[1], port[0], user[0], pw[0], port[1], "disable", "5")
	db, err := gorm.Open(config.DB.Dialect, msgArgs)

	defer db.Close()
	returnRes := models.ReturnRes{}
	if err != nil {
		fmt.Println("Failed to connect to mysql")
		returnRes.Code = "01"
		returnRes.Text = "Terjadi kendala teknis"

		response, _ := json.Marshal(returnRes)
		w.Write(response)
		return
	}

	data, _ := GetPublishArticles(db)
	var tmpl, _ = template.ParseFiles("view/index.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func GetPublishArticles(db *gorm.DB) ([]models.Article, error) {
	res := []models.Article{}
	err := db.Where("publish = 'yes'").Order("id DESC").Find(&res).Error
	if err != nil {
		fmt.Println("Failed get Articles: ", err)
		return res, err
	}
	return res, nil
}
