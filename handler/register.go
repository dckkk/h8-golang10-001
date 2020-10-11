package handler

import (
	"Golang10/Final/Ardi/models"
	"Golang10/Final/Ardi/sessions"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func GetRegister(w http.ResponseWriter, r *http.Request) {
	sessions, err := sessions.Get(r)
	if err != nil {
		fmt.Println("failed to get sessions: ", err)
		http.Error(w, "session error", 400)
		return
	}
	if sessions.Values["email"] == nil {
		http.ServeFile(w, r, "view/register.html")
	}
	http.Redirect(w, r, "/articles", 302)
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	returnRes := models.ReturnRes{}
	body, err := ioutil.ReadAll(r.Body)

	fmt.Println("request data on register : ", string(body))
	if err != nil {
		fmt.Println("Masuk returnRes")
		returnRes.Code = "422"
		returnRes.Text = "FAILED"

		response, _ := json.Marshal(returnRes)

		w.Write(response)
		return
	}

	user := &models.UserValidate{}
	if err = json.Unmarshal(body, user); err != nil {
		fmt.Println("return validator data")

		returnRes.Code = "422"
		returnRes.Text = "Gagal Tervalidasi"

		response, _ := json.Marshal(returnRes)

		w.Write(response)
		return
	}
	msg := &models.UserValidate{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if msg.Validate() == false {

		validateResponse, _ := json.Marshal(msg)

		w.Write([]byte(validateResponse))
		return
	}

	//encrypt password here
	pass, err := bcrypt.GenerateFromPassword([]byte(msg.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		returnRes.Code = "422"
		returnRes.Text = "Gagal enkripsi password"

		response, _ := json.Marshal(returnRes)
		w.Write(response)
		return
	}

	dataInsert := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(pass),
	}

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
		return
	}

	InsertNewUser(db, dataInsert)

	jsonResponse, _ := json.Marshal(user)

	fmt.Println("responses : ", string(jsonResponse))

	w.Write([]byte(jsonResponse))
	return
}

func InsertNewUser(db *gorm.DB, req models.User) error {
	err := db.Create(&req).Error
	if err != nil {
		fmt.Println("Failed insert new user: ", err)
		return err
	}
	return nil
}
