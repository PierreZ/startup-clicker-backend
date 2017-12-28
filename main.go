package main

import (
	"flag"

	"github.com/PierreZ/startup-clicker-backend/assets"
	"github.com/PierreZ/startup-clicker-backend/routes"
)

var initGTS = flag.Bool("create-gts", false, "Init GTS")

func main() {

	flag.Parse()

	assets.Init(*initGTS)
	routes.StartServer()
}
