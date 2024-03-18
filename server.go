package movierestapigo

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, router *chi.Mux) error {
	s.httpServer = &http.Server{
		Addr:			 ":" + port,
		MaxHeaderBytes:	 1 << 20, // 1 MB
		ReadTimeout:	 10 * time.Second,
		WriteTimeout:	 10 * time.Second,
		Handler: 		router,	
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.httpServer.Close()
}
