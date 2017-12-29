package main

import (
	"flag"
	"log"
	"os"

	"github.com/PierreZ/startup-clicker-backend/routes"
	"github.com/PierreZ/startup-clicker-backend/services"
)

var initGTS = flag.Bool("create-gts", false, "Init GTS")

func main() {

	flag.Parse()

	if *initGTS {
		services.CreateGTS()
		os.Exit(0)
	}

	err := services.CreateMoneyCache()
	if err != nil {
		log.Fatalln(err)
	}

	routes.StartServer()
}
