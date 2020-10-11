package handler

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/models"
	"Golang10/Final/Ardi/sessions"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

func GetAboutPost(w http.ResponseWriter, r *http.Request) {
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

	data, _ := GetAboutData(db)
	var tmpl, _ = template.ParseFiles("view/backend/about-post.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func GetAboutData(db *gorm.DB) (models.About, error) {
	res := models.About{}
	err := db.First(&res).Error
	if err != nil {
		fmt.Println("Failed get about page data: ", err)
		return res, err
	}
	return res, nil
}

func PostAboutPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if strings.TrimSpace(r.PostFormValue("title")) == "" {
		fmt.Println("error : empty title")
		response, _ := json.Marshal("Error! Judul tidak boleh kosong")
		w.Write(response)
		return
	}

	if strings.TrimSpace(r.PostFormValue("text")) == "" {
		fmt.Println("error : empty text")
		response, _ := json.Marshal("Error! Text tidak boleh kosong")
		w.Write(response)
		return
	}

	dataInsert := models.About{
		Title: r.PostFormValue("title"),
		Text:  r.PostFormValue("text"),
	}

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

	if len(data.Title) > 0 {
		UpdateAbout(db, data.ID, r.PostFormValue("title"), r.PostFormValue("text"))
		fmt.Println("Update about")
	} else {
		InsertAbout(db, dataInsert)
		fmt.Println("create about")
	}

	http.Redirect(w, r, "/about-post", 302)
}

func InsertAbout(db *gorm.DB, req models.About) error {
	err := db.Create(&req).Error
	if err != nil {
		fmt.Println("Failed insert new user: ", err)
		return err
	}
	return nil
}

func UpdateAbout(db *gorm.DB, id int, title string, text string) error {
	err := db.Exec(`UPDATE abouts SET title = ?, text = ? WHERE id = ?`, title, text, id).Error
	if err != nil {
		fmt.Println("Failed update abouts : ", err)
		return err
	}
	return nil
}
