package admin

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	Authenticate(ctx context.Context, username string) (*models.Admin, error)
}


