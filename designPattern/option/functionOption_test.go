package option

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestFunctionOption(t *testing.T) {
	addr := "127.0.0.1"
	expected := Server{
		Addr:         addr,
		Port:         defaultPort,
		Protocol:     defaultProtocol,
		ReadTimeout:  time.Duration(defaultReadTimeoutSecond) * time.Second,
		WriteTimeout: time.Duration(defaultWriteTimeoutSecond) * time.Second,
		MaxConns:     defaultMaxConns,
		TLSConfig:    nil,
		ErrorLog:     nil,
	}
	s, _ := InitServer(addr)
	assert.Equal(t, &expected, s)
}

func TestFunctionOption1(t *testing.T) {
	addr := "127.0.0.1"
	port := 8088
	protocol := "udp"
	readTimeout := time.Duration(30) * time.Second
	writeTimeout := time.Duration(60) * time.Second
	var maxConns int32 = 100
	tlsConfig := tls.Config{
		ServerName: "test",
	}
	log := log.Logger{}
	expected := Server{
		Addr:         addr,
		Port:         port,
		Protocol:     protocol,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		MaxConns:     maxConns,
		TLSConfig:    &tlsConfig,
		ErrorLog:     &log,
	}

	s, _ := InitServer(addr,
		SetPort(port), SetProtocol(protocol),
		SetReadTimeout(readTimeout), SetWriteTimeout(writeTimeout),
		SetMaxConns(maxConns), SetTLSConfig(&tlsConfig),
		SetErrorLog(&log))
	assert.Equal(t, &expected, s)
}
