package stock

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/stock"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type StockRepository struct {
	conn *sql.DB
}

func NewStockRepository(conn *sql.DB) stock.Repository {
	return &StockRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *StockRepository) getOne(ctx context.Context, query string, args ...interface{}) (*models.StockItem, error) {
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	stock := &models.StockItem{}
	err = row.Scan(&stock.ID, &stock.ClinicID, &stock.Quantity, &stock.Name, &stock.Price, &stock.LotNumber,
		&stock.ExpiryDate, &stock.ProductID, &stock.StockLeft, &stock.ImmunizationSchedule, &stock.AuthorizedInterval)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return stock, nil
}

func (repository *StockRepository) GetByID(ctx context.Context, id string) (*models.StockItem, error) {
	query := "SELECT * FROM `stock_item` WHERE id = ?"
	return repository.getOne(ctx, query, id)
}

func (repository *StockRepository) UpdateStock(ctx context.Context, value int, id string) (*sql.Tx, error) {
	tx, err := repository.conn.BeginTx(ctx, nil)
	if err != nil {
		logger.Error(err)
		return tx, err
	}

	query := `update stock_item
			set stock_left = stock_left - ?
			where id = ?`
	_, err = tx.ExecContext(ctx, query, value, id)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return tx, err
	}
	return tx, nil
}
