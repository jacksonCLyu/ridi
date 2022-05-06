package core

import (
	"github.com/jacksonCLyu/ridi-config/pkg/config"
	"github.com/jacksonCLyu/ridi-faces/pkg/env"
	"github.com/jacksonCLyu/ridi-log/log"
	"github.com/jacksonCLyu/ridi-utils/utils/errcheck"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func Boot(opts ...Option) (bootErr error) {
	defer rescueutil.Recover(func(e any) {
		bootErr = e.(error)
	})
	// env init
	errcheck.CheckAndPanic(env.Init())
	// options init
	options := DefaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	return
}

func DefaultOptions() *options {
	errcheck.CheckAndPanic(config.Init())
	errcheck.CheckAndPanic(log.Init())
	return &options{
		configer: config.L(),
		logger:   log.L(),
	}
}
