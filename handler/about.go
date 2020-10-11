package handler

import (
	"Golang10/Final/Ardi/databases"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
)

func About(w http.ResponseWriter, r *http.Request) {
	config := databases.GetConfig()
	db, err := gorm.Open(config.DB.Dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect to mysql")
		response, _ := json.Marshal("Error! Terjadi kendala teknis")
		w.Write(response)
		return
	}

	data, _ := GetAboutData(db)
	var tmpl, _ = template.ParseFiles("view/about.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
