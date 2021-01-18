package config

import (
	"net/http"
	"time"
)

func NewEchoConfig(configuration Config) *http.Server {
	return &http.Server{
		Addr:         ":" + configuration.Get("PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
}
