package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bradleyshawkins/rent/rest/person"

	"github.com/bradleyshawkins/rent/config"
	"github.com/bradleyshawkins/rent/postgres"
	"github.com/bradleyshawkins/rent/rest"
)

func main() {
	log.Println("Starting rent service")

	c, err := config.ParseConfig()
	if err != nil {
		log.Println(fmt.Errorf("unable to initialize config. Error: %v", err))
		os.Exit(0)
	}

	log.Println("Initializing database connection")
	m, err := postgres.New(c.ConnectionString, c.MigrationPath)
	if err != nil {
		log.Printf("unable to get database connection. Error: %v\n", err)
		os.Exit(1)
	}

	log.Println("Creating person router")
	pr := person.NewRouter(m)

	log.Println("Registering routes")
	router := rest.SetupRouter(pr)

	log.Println("Starting router")
	if err := router.Start(context.Background(), c.Port); err != nil {
		log.Println("unable to start router. Error:", err)
		os.Exit(2)
	}

	log.Println("Ready for traffic")
}
