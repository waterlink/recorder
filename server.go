package main

import (
	"net/http"
	"time"
)

const (
	ServerReadTimeout    = 10 * time.Second
	ServerWriteTimeout   = 10 * time.Second
	ServerMaxHeaderBytes = 1 << 20

	ServerTerminateMethod = "_TERMINATE"
)

type Server struct {
	server *http.Server
	stop   chan error
}

func NewServer(listen Listen) *Server {
	s := &Server{
		server: &http.Server{
			Addr: *listen,

			ReadTimeout:    ServerReadTimeout,
			WriteTimeout:   ServerWriteTimeout,
			MaxHeaderBytes: ServerMaxHeaderBytes,
		},
		stop: make(chan error),
	}

	s.server.Handler = http.HandlerFunc(s.Handler)

	return s
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == ServerTerminateMethod {
		s.stop <- nil
	}
}

func (s *Server) Start() error {
	go s.ListenAndServe()
	return <-s.stop
}

func (s *Server) ListenAndServe() {
	s.stop <- s.server.ListenAndServe()
}
