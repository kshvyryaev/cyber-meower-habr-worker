package pkg

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func ProvideZap() (*zap.Logger, func(), error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, nil, errors.Wrap(err, "zap")
	}

	cleanup := func() {
		logger.Sync()
	}

	return logger, cleanup, nil
}
