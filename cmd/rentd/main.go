package main

import (
	"context"
	"os"

	"github.com/bradleyshawkins/rent/kit/bhttp"

	"github.com/bradleyshawkins/rent/cmd/rentd/routes"
	"github.com/bradleyshawkins/rent/internal/platform/postgres"
	"github.com/bradleyshawkins/rent/kit/bapp/starter"
)

func main() {
	ctx := context.Background()

	db, err := postgres.New(&postgres.Config{
		Host:     "",
		Database: "",
		Schema:   "",
		Username: "",
		Password: "",
	})
	if err != nil {
		os.Exit(0)
	}

	server := bhttp.NewServer(&bhttp.Config{
		Port: 8080,
	})

	routes.Register(server, &routes.Config{
		DB: db,
	})

	err = starter.Start(ctx, server)
	if err != nil {
		os.Exit(0)
	}
}
