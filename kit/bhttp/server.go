package bhttp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bradleyshawkins/rent/kit/berror"

	"github.com/bradleyshawkins/rent/kit/bapp/starter"

	"github.com/go-chi/chi"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type Server struct {
	server http.Server
	mux    *chi.Mux
}

type Config struct {
	Port int
}

func NewServer(c *Config) *Server {
	mux := chi.NewMux()
	return &Server{
		mux: mux,
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", c.Port),
			Handler: mux,
		},
	}
}

func (s *Server) Start(ctx context.Context) (starter.Stopper, error) {

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			// TODO: Handle error
		}
	}()

	return s, nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Handle(method string, route string, h Handler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			if berr, ok := err.(*berror.Error); ok {
				berr.WriteResponse(w)
				return
			}
			berror.Internal(err, "unexpected error occurred").WriteResponse(w)
			return
		}
	}

	s.mux.Method(method, route, http.HandlerFunc(f))
}
