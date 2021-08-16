package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
)

type Router struct {
	router *chi.Mux
}

func (r *Router) Start(ctx context.Context, port string) error {
	log.Println("Starting http router...")
	srv := http.Server{
		Addr:    ":" + port,
		Handler: r.router,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		// TODO: validate that this is working as expected
		for i := range c {
			log.Println("Shutting down. Signal Interrupt", i)
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Println(fmt.Errorf("error shutting down http server. Error: %v", err))
			}
		}
	}()

	err := srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("unable to start http server. Error: %v", err)
	}

	return nil
}

type register interface {
	RegisterEndpoints(m chi.Router)
}

func SetupRouter(routers ...register) *Router {
	log.Println("Creating http router and registering endpoints...")
	c := chi.NewRouter()
	r := &Router{
		router: c,
	}

	for _, router := range routers {
		router.RegisterEndpoints(r.router)
	}

	return r
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
