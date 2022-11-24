package logger

import (
	"go.uber.org/zap"
)

func Init() (*zap.Logger, error) {
	l, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	return l, nil
}
