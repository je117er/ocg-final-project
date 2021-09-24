package stock

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*models.StockItem, error)
}
