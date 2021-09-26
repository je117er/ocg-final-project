package stock

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*models.StockItem, error)
	UpdateStock(ctx context.Context, value int, id string) (*sql.Tx, error)
}
