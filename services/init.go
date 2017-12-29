package services

import (
	"os"
	"time"

	warp "github.com/PierreZ/Warp10Exporter"
	log "github.com/sirupsen/logrus"
)

// CreateGTS is ensuring all the GTS are existing by pushing a 0 in it.
// As we are computing using bucketizer.sum all the time
func CreateGTS() {

	log.Info("Creating assets")

	for name := range All {
		gts := warp.NewGTS("asset").WithLabels(warp.Labels{"name": name}).AddDatapoint(time.Now(), 0)
		err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
		if err != nil {
			panic(err)
		}
	}

	// Creating Money
	gts := warp.NewGTS("money").AddDatapoint(time.Now(), 0)
	err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
	if err != nil {
		panic(err)
	}
}
