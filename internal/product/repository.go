package product

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	ByID(ctx context.Context, id string) (*models.Product, error)
	Fetch(ctx context.Context, query string, args ...interface{}) (*models.Product, error)
	All(ctx context.Context) ([]models.Product, error)
	AllNameID(ctx context.Context) ([]models.ProductNameID, error)
}
