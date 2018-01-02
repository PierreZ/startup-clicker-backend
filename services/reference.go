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

// All are all the assets available for startup-clicker
// Based on http://cookieclicker.wikia.com/wiki/Building
// TODO: find cool assets
var AllAssetReference = map[string]assetReference{
	"a": {
		Name:        "a",
		Description: "description de A",
		BasePrice:   15,
		Rate:        0.1,
		ID:          1,
	},
	"b": {
		Name:      "b",
		BasePrice: 100,
		Rate:      1,
		ID:        2,
	},
	"c": {
		Name:      "c",
		BasePrice: 1100,
		Rate:      8,
		ID:        3,
	},
	"d": {
		Name:      "d",
		BasePrice: 12000,
		Rate:      47,
		ID:        4,
	},
	"e": {
		Name:      "e",
		BasePrice: 130000,
		Rate:      260,
		ID:        5,
	},
	"f": {
		Name:      "f",
		BasePrice: 1400000,
		Rate:      1400,
		ID:        6,
	},
	"g": {
		Name:      "g",
		BasePrice: 20000000,
		Rate:      7800,
		ID:        7,
	},
	"h": {
		Name:      "h",
		BasePrice: 330 * 1000000,
		Rate:      44000,
		ID:        8,
	},
	"i": {
		Name:      "i",
		BasePrice: 5100 * 1000000,
		Rate:      260000,
		ID:        9,
	},
	"j": {
		Name:      "j",
		BasePrice: 75000000000,
		Rate:      1600000,
		ID:        10,
	},
	"k": {
		Name:      "k",
		BasePrice: 1000000000000,
		Rate:      10000000,
		ID:        11,
	},
	"l": {
		Name:      "l",
		BasePrice: 14000000000000,
		Rate:      65000000,
		ID:        12,
	},
	"m": {
		Name:      "m",
		BasePrice: 170000000000000,
		Rate:      430000000,
		ID:        13,
	},
	"n": {
		Name:      "n",
		BasePrice: 2100000000000000,
		Rate:      2900000000,
		ID:        14,
	},
	"o": {
		Name:      "o",
		BasePrice: 26000000000000000,
		Rate:      21000000000,
		ID:        15,
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
