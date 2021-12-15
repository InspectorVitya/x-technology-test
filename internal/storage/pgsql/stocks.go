package pgsql

import (
	"context"
	"github.com/inspectorvitya/x-technology-test/internal/model"
)

func (s *StoragePgSql) CreateStockQuotes(ctx context.Context, stocks []model.Stocks) error {
	query := `INSERT INTO stock_quotes(symbol, price, volume, last_trade)
	VALUES(:symbol, :price, :volume, :last_trade);`
	_, err := s.db.NamedExecContext(ctx, query, stocks)
	if err != nil {
		return err
	}
	return err
}
func (s *StoragePgSql) GetStockQuotes(ctx context.Context) ([]model.Stocks, error) {
	query := `SELECT symbol, price, volume, last_trade
	FROM stock_quotes;`
	var test []model.Stocks
	err := s.db.SelectContext(ctx, &test, query)
	if err != nil {
		return nil, err
	}

	return test, nil
}
func (s *StoragePgSql) UpdateStockQuotes(ctx context.Context, stocks []model.Stocks) error {
	query := `UPDATE stock_quotes
	SET price=:price, volume=:volume, last_trade=:last_trade
	WHERE symbol =:symbol;`
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	for _, val := range stocks {
		_, err := tx.NamedExec(query, val)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func (s *StoragePgSql) CheckEmptyDb(ctx context.Context) (bool, error) {
	query := `SELECT COUNT(1) FROM stock_quotes`
	row := s.db.QueryRowxContext(ctx, query)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	return true, nil
}
