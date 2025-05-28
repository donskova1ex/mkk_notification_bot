package processors

import (
	"context"
	"fmt"
)

type ClientDataRepository interface {
	ClientNumberFound(ctx context.Context, number string) (bool, error)
}

type ClientDataLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type clientData struct {
	repository ClientDataRepository
	logger     ClientDataLogger
}

func NewClientDataProcessor(repository ClientDataRepository, log ClientDataLogger) *clientData {
	return &clientData{
		repository: repository,
		logger:     log,
	}
}

func (c *clientData) ClientNumberFound(ctx context.Context, number string) (bool, error) {
	ok, err := c.repository.ClientNumberFound(ctx, number)
	if err != nil {
		return false, fmt.Errorf("error getting client number found: %w", err)
	}
	return ok, nil
}
