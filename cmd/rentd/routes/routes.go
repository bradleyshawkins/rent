package routes

import (
	"net/http"

	"github.com/bradleyshawkins/rent/kit/bhttp"

	"github.com/bradleyshawkins/rent/cmd/rentd/routes/handlers"

	"github.com/bradleyshawkins/rent/internal/user"
	"github.com/bradleyshawkins/rent/internal/user/userdb"

	"github.com/bradleyshawkins/rent/internal/platform/postgres"
)

type Config struct {
	DB *postgres.Database
}

func Register(s *bhttp.Server, c *Config) {
	userSvc := user.NewCore(userdb.NewStore(c.DB))

	userHandler := handlers.NewUser(userSvc)

	s.Handle(http.MethodPost, "/user", userHandler.CreateUser)
	s.Handle(http.MethodGet, "/user/{id}", userHandler.GetUserByID)
	s.Handle(http.MethodPut, "/user/{id}", userHandler.UpdateUser)
}
