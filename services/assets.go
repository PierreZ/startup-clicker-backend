package services

import (
	"os"

	cache "github.com/PierreZ/warpcache"
	log "github.com/sirupsen/logrus"
)

// Assets is the warpcache
var Assets *cache.MultipleCache

// CreateAssetsCache is creating a warpcache for the Assets GTS
func CreateAssetsCache() error {

	var err error

	config := cache.Configuration{
		ReadToken:  os.Getenv("RTOKEN"),
		WriteToken: os.Getenv("WTOKEN"),
		Endpoint:   os.Getenv("ENDPOINT"),
	}

	selector := cache.Selector{
		Classname: "asset",
		Labels: map[string]string{
			"name": "~.*",
		},
	}

	Assets, err = cache.NewMultipleCache(selector, "name", config)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}
