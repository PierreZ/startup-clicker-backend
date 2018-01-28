package services

import (
	"os"

	cache "github.com/PierreZ/warpcache"
	log "github.com/sirupsen/logrus"
)

type Upgrade struct {
	BasePrice float64
	Number    float64 `json:"number"`
	Price     float64 `json:"price"`
}

func GetFundRaising() Upgrade {
	upgrade := Upgrade{
		BasePrice: 50000,
		Number:    Fundraising.Get(),
		Price:     50000,
	}

	log.Printf("%+v", upgrade)
	return upgrade
}

// Fundraising is the warpcache
var Fundraising *cache.SingleCache

// CreateFundraisingCache is creating a warpcache for the money GTS
func CreateFundraisingCache() error {

	log.Info("Creating Fundraising...")

	var err error

	config := cache.Configuration{
		ReadToken:  os.Getenv("RTOKEN"),
		WriteToken: os.Getenv("WTOKEN"),
		Endpoint:   os.Getenv("ENDPOINT"),
	}

	selector := cache.Selector{
		Classname: "fundraising",
	}

	Fundraising, err = cache.NewSingleCache(selector, config)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	go watchErrors(Fundraising.Errors)

	log.Info("Creating Fundraising is OK")
	//log.Info(Fundraising.Get())

	return nil
}
