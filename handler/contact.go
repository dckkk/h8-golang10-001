package handler

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetContact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/contact.html")
}

func PostContact(w http.ResponseWriter, r *http.Request) {
	returnRes := models.ReturnRes{}
	body, err := ioutil.ReadAll(r.Body)

	fmt.Println("request data on contact form : ", string(body))
	if err != nil {
		returnRes.Code = "422"
		returnRes.Text = "Gagal Dikirim"

		response, _ := json.Marshal(returnRes)

		w.Write(response)
		return
	}

	contact := &models.Contact{}
	if err = json.Unmarshal(body, contact); err != nil {
		fmt.Println("return validator data")

		returnRes.Code = "422"
		returnRes.Text = "Gagal Tervalidasi"

		response, _ := json.Marshal(returnRes)

		w.Write(response)
		return
	}

	dataInsert := models.Contact{
		Name:    contact.Name,
		Email:   contact.Email,
		Subject: contact.Subject,
		Message: contact.Message,
	}

	config := databases.GetConfig()
	db, err := gorm.Open(config.DB.Dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect to mysql")
		returnRes.Code = "01"
		returnRes.Text = "Terjadi kendala teknis"

		response, _ := json.Marshal(returnRes)
		w.Write(response)
		return
	}

	InsertNewContact(db, dataInsert)

	returnRes.Code = "200"
	returnRes.Text = "SUCCESS"

	jsonResponse, _ := json.Marshal(returnRes)

	w.Write([]byte(jsonResponse))
	return
}

func InsertNewContact(db *gorm.DB, req models.Contact) error {
	err := db.Create(&req).Error
	if err != nil {
		fmt.Println("Failed insert new user: ", err)
		return err
	}
	return nil
}
