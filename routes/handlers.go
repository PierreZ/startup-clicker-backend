package routes

import (
	"math"
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
		return c.JSON(http.StatusBadRequest, "not possible to click more than 10 in a sec ;)")
	}

	current := services.Money.Get()

	ratePersec := 0.0
	assets := services.AllAssetReference
	for assetName, asset := range assets {
		ratePersec += services.Assets.Get(assetName) * asset.Rate
	}
	//log.Println("total:", ratePersec)
	ratePersec /= 100
	//log.Println("ratePersec /= 100:", ratePersec)

	ratePersec *= services.Fundraising.Get()
	ratePersec = math.Ceil(ratePersec)
	//log.Println("rate:", ratePersec)

	services.Money.Set(current + input.Click + input.Click*ratePersec)
	//log.Println("set:", current+input.Click*ratePersec+1)

	return c.JSON(http.StatusOK, "Click OK")
}

func getAssets(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, services.GetAssets(), "  ")
}

func getUpgrades(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, services.GetFundRaising(), "  ")
}

func buyUpgrade(c echo.Context) error {

	money := services.Money.Get()
	number := services.Fundraising.Get()
	price := 50000 * (1 + number)

	moneyLeft := money - price
	if moneyLeft < 0 {
		return c.String(http.StatusBadRequest, "Not enough money! ")
	}
	services.Fundraising.Set(number + 1)
	services.Money.Set(moneyLeft)

	return c.String(http.StatusOK, "OK")
}

func buyAsset(c echo.Context) error {
	assetName := c.Param("asset")

	// Asset exists?
	if _, ok := services.AllAssetReference[assetName]; !ok {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	number := services.Assets.Get(assetName)
	basePrice := services.AllAssetReference[assetName].BasePrice
	price := services.GetPrice(basePrice, number)
	money := services.Money.Get()

	moneyLeft := money - price
	if moneyLeft < 0 {
		return c.String(http.StatusBadRequest, "Not enough money! ")
	}

	services.Assets.Set(assetName, number+1)
	services.Money.Set(moneyLeft)

	return c.String(http.StatusOK, "OK")
}
