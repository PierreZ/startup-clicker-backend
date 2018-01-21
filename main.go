package main

import (
	"log"

	"github.com/PierreZ/startup-clicker-backend/routes"
	"github.com/PierreZ/startup-clicker-backend/services"
)

func main() {

	err := services.CreateGTS()
	if err != nil {
		log.Fatalln(err)
	}

	err = services.CreateMoneyCache()
	if err != nil {
		log.Fatalln(err)
	}

	err = services.CreateAssetsCache()
	if err != nil {
		log.Fatalln(err)
	}

	err = services.CreateFundraisingCache()
	if err != nil {
		log.Fatalln(err)
	}

	services.CreateWorkers()

	routes.StartServer()
}
