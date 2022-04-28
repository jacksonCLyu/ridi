package core

import (
	"github.com/jacksonCLyu/ridi-faces/pkg/configer"
	"github.com/jacksonCLyu/ridi-faces/pkg/logger"
)

type Option interface {
	apply(opts *options)
}

type options struct {
	configer configer.Configurable
	logger   logger.Logger
}

type applyFunc func(opts *options)

func (f applyFunc) apply(opts *options) {
	f(opts)
}

// WithConfig sets the configer.
func WithConfig(configer configer.Configurable) Option {
	return applyFunc(func(opts *options) {
		opts.configer = configer
	})
}

// WithLogger sets the logger.
func WithLogger(logger logger.Logger) Option {
	return applyFunc(func(opts *options) {
		opts.logger = logger
	})
}
