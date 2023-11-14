package todo

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Seconds,
		WriteTimeout:   10 * time.Seconds,
	}

	return s.httpServer.ListenAndServe()
}
