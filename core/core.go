package core

import (
	"github.com/jacksonCLyu/ridi-config/pkg/config"
	"github.com/jacksonCLyu/ridi-faces/pkg/env"
	"github.com/jacksonCLyu/ridi-log/log"
	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/errcheck"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func Boot(opts ...Option) (bootErr error) {
	defer rescueutil.Recover(func(e any) {
		bootErr = e.(error)
	})
	// env init
	errcheck.CheckAndPanic(env.Init())
	// default options
	options := DefaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	// global config init
	errcheck.CheckAndPanic(config.Init(config.WithConfigurable(assignutil.Assign(config.NewConfig(config.WithFilePath(options.configPath))))))
	// global log init
	errcheck.CheckAndPanic(log.Init(log.WithLogger(options.logger)))
	return
}

func DefaultOptions() *options {
	errcheck.CheckAndPanic(log.Init())
	return &options{
		logger:   log.L(),
	}
}
