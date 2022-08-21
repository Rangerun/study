package factory

import (
	"fmt"
	"sync"
	"test01/store"
)

var (
	providersMu sync.RWMutex
	providers = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store:provider is null")
	}

	if _, dup := providers[name]; dup {
		panic("twice" + name)
	}
	providers[name] = p
}



func New(ProvidersName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[ProvidersName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store:unkonw %s", ProvidersName)
	}
	return p, nil
}

