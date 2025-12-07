package http

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/reiott/vm-api/filestore"
	"go.uber.org/zap"
)

type Server struct {
	Host string `long:"host" description:"The IP to listen on" default:"0.0.0.0" env:"HOST"`
	Port int    `long:"port" description:"The port to listen on" default:"8080" env:"PORT"`

	log *zap.Logger
}

func (s *Server) Serve(context.Context) {
	s.log, _ = zap.NewDevelopment()
	defer s.log.Sync()

	service := Service{
		Store: &Store{
			VMStore: filestore.NewVMStore(),
		},
		Now: time.Now,
		Log: s.log,
	}

	handler := NewMux(MuxOpts{}, service)

	httpServer := &http.Server{
		Handler:     handler,
		IdleTimeout: 5 * time.Second,
	}

	httpServer.SetKeepAlivesEnabled(true)

	listener, err := s.NewListener()
	if err != nil {

		return
	}
	defer listener.Close()

	if err := httpServer.Serve(listener); err != nil {
		return
	}
}

func (s *Server) NewListener() (net.Listener, error) {
	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return lis, nil
}
