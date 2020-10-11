package main

import (
	"Golang10/Final/Ardi/databases"
	"Golang10/Final/Ardi/routes"
)

func main() {
	//setup database first
	config := databases.GetConfig()
	databases.SetupDatabase(config)

	//setup routes
	routes.Routes()

	//GOLANG DEVELOPED BY : ARDI GUNAWAN -> ardigunawan1992@gmail.com
}
