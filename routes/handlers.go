package routes

import (
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

	services.Money.Set(current + input.Click)

	return c.JSON(http.StatusOK, "Click OK")
}

func getAssets(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, services.GetAssets(), "  ")
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
