package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bradleyshawkins/rent/postgres"

	"github.com/bradleyshawkins/rent/config"
)

func main() {
	log.Println("Starting rent service")
	c, err := config.ParseConfig()
	if err != nil {
		log.Println(fmt.Errorf("unable to initialize config. Error: %v", err))
		os.Exit(0)
	}
	_, err = postgres.New(c.ConnectionString, c.MigrationPath)
	if err != nil {
		log.Printf("unable to get database connection. Error: %v\n", err)
		os.Exit(1)
	}
	//
	//personService := person.NewPersonService(m)
	//
	//router := http.SetupRouter(personService)
	//
	//if err := router.Start(context.Background(), ":8080"); err != nil {
	//	log.Println("unable to start router. Error:", err)
	//	os.Exit(2)
	//}

	log.Println("Ready for traffic")
}
