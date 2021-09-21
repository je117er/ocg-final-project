package admin

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context, query string, args... interface{}) (*models.Admin, error)
	ByUsername(ctx context.Context, username string) (*models.Admin, error)
}
