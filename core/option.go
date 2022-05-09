package core

import (
	"github.com/jacksonCLyu/ridi-faces/pkg/logger"
)

type Option interface {
	apply(opts *options)
}

type options struct {
	configPath string
	logger   logger.Logger
}

type applyFunc func(opts *options)

func (f applyFunc) apply(opts *options) {
	f(opts)
}

// WithConfig sets the configer.
func WithConfig(configPath string) Option {
	return applyFunc(func(opts *options) {
		opts.configPath = configPath
	})
}

// WithLogger sets the logger.
func WithLogger(logger logger.Logger) Option {
	return applyFunc(func(opts *options) {
		opts.logger = logger
	})
}
