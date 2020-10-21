package service

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// Router is the service router. The Init field is a function
// that sets the different routes/endpoints
type Router struct {
	Mux      *chi.Mux
	initFunc func(*chi.Mux)
}

// InitRoutes is a wrapper method for the InitFunc
func (r Router) InitRoutes() {
	r.initFunc(r.Mux)
}

// Service contains common elements
// for a microservice
type Service struct {
	SvcPort   string
	SvcName   string
	SvcRouter Router
}

// Start starts the service
func (s Service) Start(msg string) error {
	loadEnv()
	s.SvcPort = os.Getenv("PORT")
	s.SvcRouter.InitRoutes()
	log.Println(msg + s.SvcPort)
	return http.ListenAndServe(":"+s.SvcPort, s.SvcRouter.Mux)
}

// Builder is for building the
// Service struct with builder pattern
type Builder struct {
	s Service
}

// Build returns the built service
func (b *Builder) Build() Service {
	return b.s
}

// Name sets the service name
func (b *Builder) Name(n string) *Builder {
	b.s.SvcName = n
	return b
}

// Router sets the service router
func (b *Builder) Router(f func(*chi.Mux)) *Builder {
	b.s.SvcRouter.Mux = chi.NewRouter()
	b.s.SvcRouter.initFunc = f
	return b
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
