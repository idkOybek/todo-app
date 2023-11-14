package todo

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr = ":" + port,
		MaxHeaderBytes : 1<<20,
		Read
	}
}
