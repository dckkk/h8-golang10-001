package routes

import (
	"Golang10/Final/Ardi/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Index)

	r.HandleFunc("/about", handler.About).Methods("GET")

	r.HandleFunc("/register", handler.GetRegister).Methods("GET")
	r.HandleFunc("/register", handler.PostRegister).Methods("POST")

	r.HandleFunc("/login", handler.GetLogin).Methods("GET")
	r.HandleFunc("/login", handler.PostLogin).Methods("POST")

	r.HandleFunc("/logout", handler.Logout).Methods("GET")

	r.HandleFunc("/contact", handler.GetContact).Methods("GET")
	r.HandleFunc("/contact", handler.PostContact).Methods("POST")

	r.HandleFunc("/articles", handler.GetArticle).Methods("GET")
	r.HandleFunc("/create-article", handler.CreateArticle).Methods("GET")
	r.HandleFunc("/create-article", handler.PostCreateArticle).Methods("POST")
	r.HandleFunc("/update-article", handler.GetUpdateArticle).Methods("GET")
	r.HandleFunc("/update-article", handler.PostUpdateArticle).Methods("POST")
	r.HandleFunc("/delete-article", handler.DeleteArticle).Methods("POST")

	r.HandleFunc("/contact-list", handler.ContactList).Methods("GET")
	http.Handle("/contact-list", r)

	r.HandleFunc("/about-post", handler.GetAboutPost).Methods("GET")
	r.HandleFunc("/about-post", handler.PostAboutPost).Methods("POST")

	http.Handle("/", r)

	//theme files such as css, js, etc.
	themeRoute := http.FileServer(http.Dir("./theme/"))
	http.Handle("/theme/", http.StripPrefix("/theme/", themeRoute))

	fmt.Println("Started")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
