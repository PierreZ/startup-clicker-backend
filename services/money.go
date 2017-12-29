package services

import (
	"os"

	cache "github.com/PierreZ/warpcache"
	log "github.com/sirupsen/logrus"
)

// Money is the warpcache
var Money *cache.SingleCache

// CreateMoneyCache is creating a warpcache for the money GTS
func CreateMoneyCache() error {

	var err error

	config := cache.Configuration{
		ReadToken:  os.Getenv("RTOKEN"),
		WriteToken: os.Getenv("WTOKEN"),
		Endpoint:   os.Getenv("ENDPOINT"),
	}

	selector := cache.Selector{
		Classname: "money",
	}

	Money, err = cache.NewSingleCache(selector, config)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	go watchErrors(Money.Errors)

	return nil
}

func watchErrors(ch chan error) {
	//var err error
	for {
		err := <-ch
		log.Errorln(err)
	}
}
