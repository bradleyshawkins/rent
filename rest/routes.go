package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent/location"

	"github.com/bradleyshawkins/rent/identity"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

type Router struct {
	Router          *chi.Mux
	registrar       *identity.Registrar
	userLoader      *identity.UserLoader
	propertyCreator *location.PropertyCreator
}

func NewRouter(registrar *identity.Registrar, userLoader *identity.UserLoader, propertyCreator *location.PropertyCreator) *Router {
	log.Println("Creating Router")
	c := chi.NewRouter()

	p := &Router{
		Router:          c,
		registrar:       registrar,
		userLoader:      userLoader,
		propertyCreator: propertyCreator,
	}

	log.Println("Registering routes")
	c.Get("/health", p.Health)
	// User management
	c.Post("/users", ErrorHandler(p.RegisterUser))
	c.Get("/users/{userID}", ErrorHandler(p.LoadUser))
	c.Post("/accounts/{accountID}/users", ErrorHandler(p.RegisterUserToAccount))
	//c.Get("/user/{userID}", ErrorHandler(p.LoadUser))
	//c.Delete("/account/{accountID}/user/{userID}", ErrorHandler(p.CancelUser))
	//
	c.Post("/account/{accountID}/property", ErrorHandler(p.CreateProperty))
	//c.Delete("/account/{accountID}/property/{propertyID}", ErrorHandler(p.RemoveProperty))
	//c.Get("/account/{accountID}/property/{propertyID}", ErrorHandler(p.LoadProperty))

	return p
}

func (r *Router) Start(ctx context.Context, port string) func(ctx context.Context) error {
	srv := http.Server{
		Addr:    ":" + port,
		Handler: r.Router,
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
