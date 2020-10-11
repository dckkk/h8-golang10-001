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

func GetArticle(w http.ResponseWriter, r *http.Request) {
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

	data, _ := GetArticles(db)
	var tmpl, _ = template.ParseFiles("view/backend/article.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func GetArticles(db *gorm.DB) ([]models.Article, error) {
	res := []models.Article{}
	err := db.Order("id DESC").Find(&res).Error
	if err != nil {
		fmt.Println("Failed get Articles: ", err)
		return res, err
	}
	return res, nil
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
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

	http.ServeFile(w, r, "view/backend/article-post.html")
}

func PostCreateArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if strings.TrimSpace(r.PostFormValue("title")) == "" {
		fmt.Println("error : empty article title")
		response, _ := json.Marshal("Error! Judul tidak boleh kosong")
		w.Write(response)
		return
	}

	if strings.TrimSpace(r.PostFormValue("text")) == "" {
		fmt.Println("error : empty article text")
		response, _ := json.Marshal("Error! Text tidak boleh kosong")
		w.Write(response)
		return
	}

	dataInsert := models.Article{
		Title:   r.PostFormValue("title"),
		Text:    r.PostFormValue("text"),
		Publish: r.PostFormValue("publish"),
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

	InsertArticle(db, dataInsert)
	fmt.Println("finish create article")

	http.Redirect(w, r, "/articles", 302)
}

func InsertArticle(db *gorm.DB, req models.Article) error {
	err := db.Create(&req).Error
	if err != nil {
		fmt.Println("Failed insert new user: ", err)
		return err
	}
	return nil
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if strings.TrimSpace(r.PostFormValue("id")) == "" {
		fmt.Println("error : empty id article")
		response, _ := json.Marshal("Error! Artikel tidak tersedia")
		w.Write(response)
		return
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

	DeleteRecord(db, r.PostFormValue("id"))
	fmt.Println("finish delete article id : ", r.PostFormValue("id"))

	http.Redirect(w, r, "/articles", 302)
}

func DeleteRecord(db *gorm.DB, id string) error {
	article := models.Article{}
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		fmt.Println("Failed to delete article: ", err)
		return err
	}
	return nil
}

func GetUpdateArticle(w http.ResponseWriter, r *http.Request) {
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

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}
	// Query()["id"] will return an array of items,
	// we only want the single item.
	id := ids[0]
	fmt.Println("Url Param 'id' is: " + string(id))

	if strings.TrimSpace(id) == "" {
		fmt.Println("error : empty id article")
		response, _ := json.Marshal("Error! Artikel tidak tersedia")
		w.Write(response)
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

	data, _ := GetArticleId(db, string(id))
	var tmpl, _ = template.ParseFiles("view/backend/article-update.html")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func GetArticleId(db *gorm.DB, id string) (models.Article, error) {
	res := models.Article{}
	err := db.Where("id = ?", id).First(&res).Error
	if err != nil {
		fmt.Println("Failed get article by id. error: ", err)
		return res, err
	}
	return res, nil
}

func PostUpdateArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if strings.TrimSpace(r.PostFormValue("title")) == "" || strings.TrimSpace(r.PostFormValue("text")) == "" {
		fmt.Println("error : empty article title")
		response, _ := json.Marshal("Error! Judul dan text tidak boleh kosong")
		w.Write(response)
		return
	}

	if strings.TrimSpace(r.PostFormValue("id")) == "" {
		fmt.Println("error : empty article text")
		response, _ := json.Marshal("Error! Artikel tidak ditemukan")
		w.Write(response)
		return
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

	data, _ := GetArticleId(db, r.PostFormValue("id"))

	if len(data.Title) > 0 {
		UpdateArticle(db, data.ID, r.PostFormValue("title"), r.PostFormValue("text"), r.PostFormValue("publish"))
		fmt.Println("finish Update article id : ", data.ID)
	}

	http.Redirect(w, r, "/articles?id="+r.PostFormValue("id"), 302)
}

func UpdateArticle(db *gorm.DB, id int, title string, text string, publish string) error {
	err := db.Exec(`UPDATE articles SET title = ?, text = ?, publish = ? WHERE id = ?`, title, text, publish, id).Error
	if err != nil {
		fmt.Println("Failed update article : ", err)
		return err
	}
	return nil
}
