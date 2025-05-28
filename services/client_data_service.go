package services

import (
	"context"
	"fmt"
	"log/slog"
)

type ClientsDataProcessor interface {
	ClientNumberFound(ctx context.Context, number string) (bool, error)
}

type ClientDataService struct {
	clientDataProcessor ClientsDataProcessor
	logger              *slog.Logger
}

func NewClientDataService(clientDataProcessor ClientsDataProcessor, logger *slog.Logger) *ClientDataService {
	return &ClientDataService{
		clientDataProcessor: clientDataProcessor,
		logger:              logger,
	}
}

func (clientDataService *ClientDataService) ClientNumberFound(ctx context.Context, number string) (bool, error) {
	c, err := clientDataService.clientDataProcessor.ClientNumberFound(ctx, number)
	if err != nil {
		return false, fmt.Errorf("error getting client number found: %w", err)
	}
	return c, nil
}
