package handler

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/models"
	"Golang10/Final/Ardi/sessions"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
)

func ContactList(w http.ResponseWriter, r *http.Request) {
	sessions, err := sessions.Get(r)
	if err != nil {
		fmt.Println("failed to get sessions: ", err)
		http.Redirect(w, r, "/login", 302)
	}
	if sessions.Values["email"] == nil {
		fmt.Println("failed to get sessions: ", err)
		http.Redirect(w, r, "/login", 302)
	}
	err = sessions.Save(r, w)
	if err != nil {
		fmt.Println("failed to save sessions: ", err)
		http.Redirect(w, r, "/login", 302)
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

	data, _ := GetContactLists(db)
	var tmpl, _ = template.ParseFiles("view/backend/contact-list.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func GetContactLists(db *gorm.DB) ([]models.Contact, error) {
	res := []models.Contact{}
	err := db.Order("id DESC").Find(&res).Error
	if err != nil {
		fmt.Println("Failed get user: ", err)
		return res, err
	}
	return res, nil
}
