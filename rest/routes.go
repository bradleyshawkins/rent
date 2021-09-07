package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

type Router struct {
	router *chi.Mux
	ps     rent.PersonStore
}

func NewRouter(ps rent.PersonStore) *Router {
	log.Println("Creating router")
	c := chi.NewRouter()

	p := &Router{
		router: c,
		ps:     ps,
	}

	log.Println("Registering person routes")
	c.Get("/health", p.Health)
	// Person management
	c.Post("/person/register", ErrorHandler(p.RegisterPerson))
	c.Get("/person/{personID}", ErrorHandler(p.LoadPerson))

	return p
}

func (r *Router) Start(ctx context.Context, port string) func(ctx context.Context) error {
	srv := http.Server{
		Addr:    ":" + port,
		Handler: r.router,
	}

	go func() {
		log.Println("Starting http server ...")
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Error shutting down server. Error:", err)
		}
	}()

	return func(ctx context.Context) error {
		log.Println("Shutting down http server ...")
		return srv.Shutdown(ctx)
	}
}

// person
//

// Landlord
// 	Properties
// 		Add a property to their list of properties
// 		Remove a property from their list of properties
// 		Edit a property from their list of properties
// 		Get a list of their properties
// 		Get details of their property
// 	Tenants
//		Approve tenant for property
//		Remove tenant from property??
//		View all tenants in owned properties
//		View tenant/s in owned property

// Tenant
//	Properties
//		Apply for property
//
