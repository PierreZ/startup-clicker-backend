package services

import (
	"math"
)

type assetReference struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BasePrice   float64
	Price       float64 `json:"price"`
	Rate        float64 `json:"rate"`
	ID          int
}

// AllAssetReference are all the assets available for startup-clicker
// Based on http://cookieclicker.wikia.com/wiki/Building
// TODO: find cool assets
var AllAssetReference = map[string]assetReference{
	"grandma": {
		Name:        "Grandmas",
		Description: "Your grand-mother and her friends do not underestand what you are doing, but they are giving you some coins to help you.",
		BasePrice:   15,
		Rate:        1,
		ID:          1,
	},
	"teammate": {
		Name:        "Former startup week-end teammate",
		Description: "Former teammate from the Startup Weekend you intended are joining your team.",
		BasePrice:   100,
		Rate:        5,
		ID:          2,
	},
	"interns": {
		Name:        "Interns",
		Description: "Why should you pay them?",
		BasePrice:   1100,
		Rate:        15,
		ID:          3,
	},
	"junior": {
		Name:        "Junior Engineer",
		Description: "They are not too expensive, and can be replaced if broken.",
		BasePrice:   12000,
		Rate:        47,
		ID:          4,
	},
	"sales": {
		Name:        "Sales team",
		Description: "To start selling a product, you need sales people.",
		BasePrice:   130000,
		Rate:        260,
		ID:          5,
	},
	"product": {
		Name:        "New product",
		Description: "By creating a new product, you are invading another market.",
		BasePrice:   1400000,
		Rate:        1400,
		ID:          6,
	},
	"office": {
		Name:        "New office",
		Description: "Opening new offices around the world will help you get worldwide",
		BasePrice:   20000000,
		Rate:        7800,
		ID:          7,
	},
	"manager": {
		Name:        "Managers",
		Description: "Someone has to look after the interns and the juniors, right?",
		BasePrice:   330 * 1000000,
		Rate:        44000,
		ID:          8,
	},
	"consultant": {
		Name:        "External consultant",
		Description: "They are cheaper, that's all that matters.",
		BasePrice:   5100 * 1000000,
		Rate:        260000,
		ID:          9,
	},
	"orga": {
		Name:        "Changing internal organizations",
		Description: "The consultant you bought earlier told you to do so.",
		BasePrice:   75000000000,
		Rate:        1600000,
		ID:          10,
	},
	"competing": {
		Name:        "Buying competing company",
		Description: "It's an easy move to dominate the market.",
		BasePrice:   1000000000000,
		Rate:        10000000,
		ID:          11,
	},
	"cryptocurrency": {
		Name:        "Cryptocurrency",
		Description: "By designing your own currency, you do not care about the dollar index",
		BasePrice:   14000000000000,
		Rate:        65000000,
		ID:          12,
	},
	"serverless": {
		Name:        "Serverless servers",
		Description: "One of the consultant told you to reduce IT costs.",
		BasePrice:   170000000000000,
		Rate:        430000000,
		ID:          13,
	},
	"planetary": {
		Name:        "New planetary office",
		Description: "One of the consultant told you to open offices on other planets to reach intergalactic market.",
		BasePrice:   2100000000000000,
		Rate:        2900000000,
		ID:          14,
	},
	"tardis": {
		Name:        "Time machine",
		Description: "By going back in time, you can sell your product throughout the history of mankind.",
		BasePrice:   26000000000000000,
		Rate:        21000000000,
		ID:          15,
	},
}

type Asset struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Number      float64 `json:"number"`
	Rate        float64 `json:"rate"`
	ID          int     `json:"id"`
}

type AssetsList []Asset

// GetAssets is returning all the assets
func GetAssets() AssetsList {
	response := AssetsList{}
	for _, asset := range AllAssetReference {
		response = append(response, Asset{
			ID:          asset.ID,
			Name:        asset.Name,
			Description: asset.Description,
			Number:      Assets.Get(asset.Name),
			Price:       GetPrice(asset.BasePrice, Assets.Get(asset.Name)),
			Rate:        asset.Rate,
		})
	}

	return response
}

// GetPrice returns the price
func GetPrice(basePrice, number float64) float64 {
	return math.Ceil(basePrice * (math.Pow(1.15, number)))
}
