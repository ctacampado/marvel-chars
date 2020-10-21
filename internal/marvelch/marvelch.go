package marvelch

import "marvel-chars/internal/service"

// MarvelCharSvc service
type MarvelCharSvc struct {
	Svc service.Service
}

// Start the marvel characters service
func (m MarvelCharSvc) Start(msg string) error {
	b := service.Builder{}
	m.Svc = b.
		Name("characters").
		Router(SetRoutes).
		Build()
	return m.Svc.Start(msg)
}
