package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

func About(w http.ResponseWriter, r *http.Request) {
	configDb := os.Getenv("DATABASE_URL")
	configDb = strings.ReplaceAll(configDb, "postgres://", "")
	user := strings.Split(configDb, ":")
	pw := strings.Split(user[1], "@")
	port := strings.Split(user[2], "/")
	msgArgs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", pw[1], port[0], user[0], pw[0], port[1], "disable", "5")
	db, err := gorm.Open(config.DB.Dialect, msgArgs)

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
