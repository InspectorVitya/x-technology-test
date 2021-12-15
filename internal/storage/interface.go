package storage

import (
	"context"
	"github.com/inspectorvitya/x-technology-test/internal/model"
)

type StorageStocks interface {
	CreateStockQuotes(ctx context.Context, stocks []model.Stocks) error
	GetStockQuotes(ctx context.Context) ([]model.Stocks, error)
	UpdateStockQuotes(ctx context.Context, stocks []model.Stocks) error
	CheckEmptyDb(ctx context.Context) (bool, error)
}
