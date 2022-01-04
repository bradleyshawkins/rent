package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bradleyshawkins/rent/location"

	"github.com/bradleyshawkins/rent/identity"

	"github.com/bradleyshawkins/rent/config"
	"github.com/bradleyshawkins/rent/postgres"
	"github.com/bradleyshawkins/rent/rest"
)

func main() {
	log.Println("Starting rent service")
	ctx := context.Background()

	c, err := config.ParseConfig()
	if err != nil {
		log.Println(fmt.Errorf("unable to initialize config. Error: %v", err))
		os.Exit(0)
	}

	log.Println("Initializing database connection")
	db, err := postgres.NewDatabase(c.ConnectionString)
	if err != nil {
		log.Println("unable to connect to database. Error:", err)
		os.Exit(1)
	}

	err = db.Migrate(c.MigrationPath)
	if err != nil {
		log.Println("unable to migrate database. Error:", err)
		os.Exit(2)
	}

	personRegistrationService := identity.NewRegistrar(db)
	propertyCreator := location.NewCreator(db)

	router := rest.NewRouter(personRegistrationService, propertyCreator)

	stop := router.Start(context.Background(), c.Port)
	if err != nil {
		log.Println("unable to start router. Error:", err)
		os.Exit(3)
	}

	log.Println("Ready for traffic")

	if err := waitForShutdown(ctx, stop); err != nil {
		log.Println("Error shutting down. Error:", err)
	}

}

func waitForShutdown(ctx context.Context, stopFunc func(ctx context.Context) error) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case <-c:
	}

	if err := stopFunc(ctx); err != nil {
		return err
	}
	return nil
}
