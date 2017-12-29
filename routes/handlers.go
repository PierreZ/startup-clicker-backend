package routes

import (
	"log"
	"net/http"

	"github.com/PierreZ/startup-clicker-backend/services"
	"github.com/labstack/echo"
)

type inputAddMoney struct {
	Click float64 `json:"click" form:"click" query:"click"`
}

func addMoney(c echo.Context) error {
	var err error
	var input inputAddMoney
	if err = c.Bind(&input); err != nil {
		return err
	}

	if input.Click > 10.0 {
		return c.JSON(http.StatusBadRequest, "not possible to click more than 10 in a sec")
	}

	current := services.Money.Get()

	log.Println("current:", current)

	services.Money.Set(current + input.Click)

	return c.JSON(http.StatusOK, "Click OK")
}

func getAssets(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, services.All, "  ")
}

// func buyAsset(c echo.Context) error {
// 	assetName := c.Param("asset")

// 	// Asset exists?
// 	if _, ok := services.All[assetName]; !ok {
// 		return c.String(http.StatusBadRequest, "Bad request")
// 	}

// 	asset := assets.GetAsset(assetName)
// 	money := assets.GetMoney()

// 	price := assets.GetPrice(asset)
// 	money = money - price
// 	if money < 0 {
// 		return c.String(http.StatusBadRequest, "Not enough money! ")
// 	}

// 	assets.SetMoney(money)

// 	gts := warp.NewGTS("asset").WithLabels(warp.Labels{"name": assetName}).AddDatapoint(time.Now(), 1)
// 	err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, "Bad request")
// 	}
// 	return c.String(http.StatusOK, "OK")
// }
