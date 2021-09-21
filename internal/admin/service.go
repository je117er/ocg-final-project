package admin

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	ByUsername(ctx context.Context, username string) (*models.Admin, error)
}
