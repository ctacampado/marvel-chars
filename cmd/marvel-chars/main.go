package main

import (
	"log"
	"marvel-chars/internal/service"

	"github.com/go-chi/chi"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b := service.Builder{}
	s := b.
		Name("characters").
		Router(chi.NewRouter(), SetRoutes).
		Build()
	return s.Start("listening at localhost:")
}
