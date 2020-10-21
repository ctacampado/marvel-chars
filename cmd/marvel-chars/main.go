package main

import (
	"log"
	"marvel-chars/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b := service.Builder{}
	s := b.
		Name("characters").
		Router(SetRoutes).
		Build()
	return s.Start("listening at localhost:")
}
