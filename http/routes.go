package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bradleyshawkins/rent"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

type personService interface {
	GetPerson(id uuid.UUID) (*rent.Person, error)
	Register(p *rent.Person) (uuid.UUID, error)
	UpdatePerson(p *rent.Person) error
	DeletePerson(id uuid.UUID) error
}

type Router struct {
	router        *chi.Mux
	personService personService
}

func (r *Router) Start(ctx context.Context, addr string) error {
	log.Println("Starting http router...")
	srv := http.Server{
		Addr:    addr,
		Handler: r.router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("unable to start http server. Error: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		// TODO: Validate that this is working as expected
		for i := range c {
			log.Println("Shutting down. Signal Interrupt", i)
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Println(fmt.Errorf("error shutting down http server. Error: %v", err))
			}
		}
	}()

	return nil
}

func SetupRouter(p personService) *Router {
	log.Println("Creating http router and registering endpoints...")
	c := chi.NewRouter()
	r := &Router{
		router:        c,
		personService: p,
	}

	r.registerEndpoints()

	return r
}

func (r *Router) registerEndpoints() {

	r.router.Route("/Person", func(router chi.Router) {
		router.Get("/{personID}", r.GetPerson)
		router.Put("/{personID}", r.UpdatePerson)
		router.Delete("/{personID}", r.DeletePerson)
	})

	r.router.Post("/register", r.Register)
}
