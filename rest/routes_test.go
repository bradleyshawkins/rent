package rest_test

import (
	"flag"
	"github.com/bradleyshawkins/rent/identity"
	"github.com/bradleyshawkins/rent/location"
	"github.com/bradleyshawkins/rent/postgres"
	"github.com/bradleyshawkins/rent/rest"
	"log"
	"os"
	"testing"
)

var (
	router *rest.Router
)

func TestMain(m *testing.M) {
	// flag.Parse() must be called before testing.Short() or else it will panic
	flag.Parse()
	// Check to see if -short argument was used on go test to signify not to run integration tests
	if testing.Short() {
		log.Println("Skipping Integration Tests")
		os.Exit(0)
	}
	// Create Database
	log.Println("Beginning integration tests")
	db, err := postgres.NewDatabase("postgresql://postgres:password@localhost:5432/rent?sslmode=disable", "../postgres/schema")
	if err != nil {
		log.Println("Unable to create database connection. Error:", err)
		os.Exit(999)
	}
	// Create registrar
	r := identity.NewRegistrar(db)
	// Create location creator
	l := location.NewCreator(db)
	// Create Router
	router = rest.NewRouter(r, l)

	code := m.Run()

	log.Println("Completed integration tests")
	os.Exit(code)
}
