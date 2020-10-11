package main

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/routes"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	//setup database first
	config := databases.GetConfig()
	databases.SetupDatabase(config)

	// migrate db
	m, err := migrate.New(
		"golang.sql",
		os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	//setup routes
	routes.Routes()

	//GOLANG DEVELOPED BY : ARDI GUNAWAN -> ardigunawan1992@gmail.com
}
