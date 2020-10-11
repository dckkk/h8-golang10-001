package handler

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	config := databases.GetConfig()
	db, err := gorm.Open(config.DB.Dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

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
