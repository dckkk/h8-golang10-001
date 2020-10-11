package handler

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/models"
	"Golang10/Final/Ardi/sessions"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	config := databases.GetConfig()
	db, err := gorm.Open(config.DB.Dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

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
