package main

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/routes"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	//setup database first
	config := databases.GetConfig()
	databases.SetupDatabase(config)

	// migrate db
	m, err := migrate.New(
		"",
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
