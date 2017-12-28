package routes

import (
	"net/http"
	"os"
	"time"

	warp "github.com/PierreZ/Warp10Exporter"
	assets "github.com/PierreZ/startup-clicker-backend/assets"
	"github.com/labstack/echo"
)

func getAssets(c echo.Context) error {

	return c.JSONPretty(http.StatusOK, assets.GetAll(), "  ")
}

func buyAsset(c echo.Context) error {
	assetName := c.Param("asset")

	// Asset exists?
	if _, ok := assets.All[assetName]; !ok {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	asset := assets.GetAsset(assetName)
	money := assets.GetMoney()

	price := assets.GetPrice(asset)
	money = money - price
	if money < 0 {
		return c.String(http.StatusBadRequest, "Not enough money! ")
	}

	assets.SetMoney(money)

	gts := warp.NewGTS("asset").WithLabels(warp.Labels{"name": assetName}).AddDatapoint(time.Now(), 1)
	err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	return c.String(http.StatusOK, "OK")
}
