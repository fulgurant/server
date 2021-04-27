package server

import (
	"github.com/fulgurant/health"
	"go.uber.org/zap"
)

type Options struct {
	config *Config
	logger *zap.Logger
	health health.IHealth
}

func DefaultOptions() *Options {
	return &Options{}
}

// WithConfig sets up stuff from the config
func (o *Options) WithConfig(value *Config) *Options {
	o.config = value
	return o
}

// WithHealth sets the health checker used to allow health checks
// to fail before actually refusing to accept new connections
func (o *Options) WithHealth(value health.IHealth) *Options {
	o.health = value
	return o
}

// WithLogger sets the logger for the server
func (o *Options) WithLogger(value *zap.Logger) *Options {
	o.logger = value.With(zap.String("module", "server"))
	return o
}
