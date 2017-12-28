package assets

import (
	"math"
	"sync"
)

type asset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BasePrice   float64
	Price       float64 `json:"price"`
	Number      float64 `json:"number"`
	Rate        float64 `json:"rate"`
}

type assetMap struct {
	v   map[string]asset
	mux sync.Mutex
}

var assets assetMap

// GetAssetNumber returns the current value of the counter for the given key.
func GetAssetNumber(asset string) float64 {
	assets.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer assets.mux.Unlock()
	a, ok := assets.v[asset]
	if ok {
		return a.Number
	}
	return -1
}

// GetAssetNumber returns the current value of the counter for the given key.
func GetAsset(asset string) asset {
	assets.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer assets.mux.Unlock()
	return assets.v[asset]
}

func GetAll() map[string]asset {
	assets.mux.Lock()
	defer assets.mux.Unlock()
	return assets.v
}

// SetAssetNumber returns the current value of the counter for the given key.
func SetAssetNumber(asset string, f float64) float64 {
	assets.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer assets.mux.Unlock()
	a, ok := assets.v[asset]
	if ok {
		a.Number = f
		a.Price = GetPrice(a)
		assets.v[asset] = a
	}
	return 0
}

func SetAsset(key string, value asset) {
	assets.mux.Lock()
	value.Price = GetPrice(value)
	assets.v[key] = value
	defer assets.mux.Unlock()
}

// Inc increment an asset from a number
func Inc(asset string) float64 {
	assets.mux.Lock()
	defer assets.mux.Unlock()

	a, ok := assets.v[asset]
	if ok {
		a.Number = a.Number + 1
		a.Price = GetPrice(a)
		assets.v[asset] = a
	}
	return 0
}

func GetPrice(a asset) float64 {
	return a.BasePrice * (math.Pow(1.15, a.Number))
}

// Init is getting the assets
func Init(init bool) {

	if init {
		createGTS()
	}

	// Create sync map
	assets.v = make(map[string]asset)
	for key, value := range All {
		SetAsset(key, value)
	}

	// Get assets numbers
	refreshAssets()

	refreshMoney()

	// Watch GTS
	go watch()

	go Generate()
}
