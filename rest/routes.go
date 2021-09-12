package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

type Router struct {
	router    *chi.Mux
	ps        rent.PersonStore
	propStore rent.PropertyStore
}

func NewRouter(ps rent.PersonStore, propStore rent.PropertyStore) *Router {
	log.Println("Creating router")
	c := chi.NewRouter()

	p := &Router{
		router:    c,
		ps:        ps,
		propStore: propStore,
	}

	log.Println("Registering person routes")
	c.Get("/health", p.Health)
	// Person management
	c.Post("/person/register", ErrorHandler(p.RegisterPerson))
	c.Get("/person/{personID}", ErrorHandler(p.LoadPerson))
	c.Delete("/account/{accountID}/person/{personID}", ErrorHandler(p.CancelPerson))

	c.Post("/account/{accountID}/property", ErrorHandler(p.RegisterProperty))
	c.Delete("/account/{accountID}/property/{propertyID}", ErrorHandler(p.RemoveProperty))

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

func getURLParamAsUUID(r *http.Request, paramName string) (uuid.UUID, error) {
	perID := chi.URLParam(r, paramName)
	if perID == "" {
		return uuid.UUID{}, rent.NewError(fmt.Errorf("%s is required", paramName), rent.WithInvalidFields(rent.InvalidField{
			Field:  paramName,
			Reason: rent.ReasonMissing,
		}), rent.WithMessage(fmt.Sprintf("%s is a required field", paramName)))
	}

	pID, err := uuid.FromString(perID)
	if err != nil {
		return uuid.UUID{}, rent.NewError(err, rent.WithInvalidFields(rent.InvalidField{
			Field:  paramName,
			Reason: rent.ReasonInvalid,
		}), rent.WithMessage(fmt.Sprintf("%s must be a UUID", paramName)))
	}
	return pID, nil
}
