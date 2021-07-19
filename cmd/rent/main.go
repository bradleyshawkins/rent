package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bradleyshawkins/rent/config"
	"github.com/bradleyshawkins/rent/postgres"
	"github.com/bradleyshawkins/rent/rest"
	"github.com/bradleyshawkins/rent/rest/person"
)

func main() {
	log.Println("Starting rent service")

	c, err := config.ParseConfig()
	if err != nil {
		log.Println(fmt.Errorf("unable to initialize config. Error: %v", err))
		os.Exit(0)
	}

	m, err := postgres.New(c.ConnectionString, c.MigrationPath)
	if err != nil {
		log.Printf("unable to get database connection. Error: %v\n", err)
		os.Exit(1)
	}

	landlordRouter := person.NewPersonRouter(m)

	router := rest.SetupRouter(landlordRouter)

	if err := router.Start(context.Background(), c.Port); err != nil {
		log.Println("unable to start router. Error:", err)
		os.Exit(2)
	}

	log.Println("Ready for traffic")
}
