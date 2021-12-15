package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/inspectorvitya/x-technology-test/internal/application"
	"github.com/inspectorvitya/x-technology-test/internal/config"
	"net"
	"net/http"
	_ "net/http/pprof"
)

type Server struct {
	router     *mux.Router
	HTTPServer *http.Server
	App        *application.App
}

func New(cfg config.Config, App *application.App) *Server {
	router := mux.NewRouter()
	server := &Server{
		HTTPServer: &http.Server{
			Addr:    net.JoinHostPort("", cfg.PortHTTP),
			Handler: router,
		},
		router: router,
		App:    App,
	}

	return server
}

func (s *Server) Start() error {
	s.router.HandleFunc("/", s.GetStocks).Methods(http.MethodGet)
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
