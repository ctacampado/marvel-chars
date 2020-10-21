package marvelch

import (
	"marvel-chars/internal/cache"
	"marvel-chars/internal/service"
)

// MarvelCharSvc service
type MarvelCharSvc struct {
	Svc   service.Service
	Cache cache.Store
}

// Start the marvel characters service
func (m *MarvelCharSvc) Start(msg string) error {
	b := service.Builder{}
	m.Svc = b.
		Name("characters").
		Router(SetRoutes).
		Build()

	err := initLoadCache(&m.Cache)
	if err != nil {
		return err
	}

	return m.Svc.Start(msg)
}

// InitLoadCache loads the cache with initial data from Marvel API.
func initLoadCache(c *cache.Store) error {
	return nil
}
