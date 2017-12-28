package assets

import "sync"

// money is safe to use concurrently.
type money struct {
	v   float64
	mux sync.Mutex
}

func (m *money) Get() float64 {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.v
}

func (m *money) Set(f float64) {
	m.mux.Lock()
	m.v = f
	m.mux.Unlock()
}

var account money

func GetMoney() float64 {
	return account.Get()
}

func SetMoney(f float64) {
	account.Set(f)
}
