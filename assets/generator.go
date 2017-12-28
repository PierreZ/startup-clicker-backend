package assets

import (
	"os"
	"time"

	warp "github.com/PierreZ/Warp10Exporter"
)

func Generate() {

	for key, _ := range GetAll() {
		go NewGenerator(key)
	}
}

func NewGenerator(name string) {
	rate := GetAsset(name).Rate

	ticker := time.NewTicker(time.Second * 1)
	for {
		for _ = range ticker.C {
			n := GetAssetNumber(name)
			if n == 0 {
				continue
			}
			money := n * rate
			money = money + account.Get()
			// Creating Money
			account.Set(money)
			gts := warp.NewGTS("money").AddDatapoint(time.Now(), money)
			err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
			if err != nil {
				panic(err)
			}
		}
	}
}
