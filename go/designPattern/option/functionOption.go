package option

import (
	"crypto/tls"
	"errors"
	"log"
	"time"
)

/*
	参考：https://coolshell.cn/articles/21146.html
*/

const (
	defaultPort               = 8086
	defaultReadTimeoutSecond  = 15
	defaultWriteTimeoutSecond = 30
	defaultMaxConns           = 10
	defaultProtocol           = "tcp"
)

type Server struct {
	Addr         string
	Port         int
	Protocol     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxConns     int32
	TLSConfig    *tls.Config
	ErrorLog     *log.Logger
}

func InitServer(addr string, options ...func(*Server)) (*Server, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be empty")
	}
	srv := Server{
		Addr:         addr,
		Port:         defaultPort,
		Protocol:     defaultProtocol,
		ReadTimeout:  time.Duration(defaultReadTimeoutSecond) * time.Second,
		WriteTimeout: time.Duration(defaultWriteTimeoutSecond) * time.Second,
		MaxConns:     defaultMaxConns,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

type Option func(server *Server)

func SetPort(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}

func SetProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func SetReadTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = t
	}
}

func SetWriteTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = t
	}
}

func SetMaxConns(conn int32) Option {
	return func(s *Server) {
		s.MaxConns = conn
	}
}

func SetTLSConfig(cfg *tls.Config) Option {
	return func(s *Server) {
		s.TLSConfig = cfg
	}
}

func SetErrorLog(log *log.Logger) Option {
	return func(s *Server) {
		s.ErrorLog = log
	}
}
