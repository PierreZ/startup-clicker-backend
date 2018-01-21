package services

import (
	"time"
)

// CreateWorkers is creating the workers to generate money
func CreateWorkers() {
	for assetName, asset := range AllAssetReference {
		go newWorker(assetName, asset.Rate)
	}

}

func newWorker(name string, rate float64) {

	ticker := time.NewTicker(time.Second * 1)
	for {
		for _ = range ticker.C {
			n := Assets.Get(name)
			if n == 0 {
				continue
			}
			money := n * rate
			money = money + Money.Get()
			// Creating Money
			Money.Set(money)

		}
	}
}
