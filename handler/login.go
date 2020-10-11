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

// var store = sessions.NewCookieStore([]byte("random"))

func GetLogin(w http.ResponseWriter, r *http.Request) {
	sessions, err := sessions.Get(r)
	if err != nil {
		fmt.Println("failed to get sessions: ", err)
		http.Error(w, "session error", 400)
		return
	}
	if sessions.Values["email"] == nil {
		http.ServeFile(w, r, "view/login.html")
	}
	http.Redirect(w, r, "/articles", 302)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	loginRes := models.ReturnRes{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("There`s error: ", err)
		loginRes.Code = "01"
		loginRes.Text = "Gagal"
		response, _ := json.Marshal(loginRes)

		w.Write(response)
		return
	}

	loginReq := models.LoginRequest{}
	if err = json.Unmarshal(body, &loginReq); err != nil {
		fmt.Println("failed to parse data: ", err)
		loginRes.Code = "01"
		loginRes.Text = "Invalid request"
		response, _ := json.Marshal(loginRes)

		w.Write(response)
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
	if err != nil {
		fmt.Println("Failed to connect to mysql on login handler")
	}

	user := &models.User{}

	if err := db.Where("email = ?", loginReq.Email).First(user).Error; err != nil {
		loginRes.Code = "01"
		loginRes.Text = "Email address not found"

		response, _ := json.Marshal(loginRes)
		w.Write(response)
		return
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println("failed to get sessions: ", err)
		loginRes.Code = "01"
		loginRes.Text = "Invalid username or password. Please try again"

		response, _ := json.Marshal(loginRes)
		w.Write(response)
		return
	}

	sessions, err := sessions.Get(r)
	if err != nil {
		fmt.Println("failed to get sessions: ", err)
		loginRes.Code = "01"
		loginRes.Text = "Sesi Gagal"

		response, _ := json.Marshal(loginRes)
		w.Write(response)
		return
	}
	sessions.Values["email"] = loginReq.Email
	err = sessions.Save(r, w)
	if err != nil {
		fmt.Println("failed to save sessions: ", err)
		loginRes.Code = "01"
		loginRes.Text = "Gagal"

		response, _ := json.Marshal(loginRes)
		w.Write(response)
		return
	}

	loginRes.Code = "00"
	loginRes.Text = "Sukses"
	response, _ := json.Marshal(loginRes)

	fmt.Println("success save sessions: ", err)

	w.Write(response)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessions, err := sessions.Get(r)
	if err != nil {
		fmt.Println("failed to get sessions: ", err)
		http.Error(w, "session error", 400)
		return
	}
	sessions.Options.MaxAge = -1
	if err := sessions.Save(r, w); err != nil {
		fmt.Println("failed to save sessions: ", err)
		http.Error(w, "session error", 400)
		return
	}

	fmt.Println("Logout from session")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
