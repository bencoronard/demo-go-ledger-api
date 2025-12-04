package config

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
