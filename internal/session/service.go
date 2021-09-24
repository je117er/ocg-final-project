package session

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetAllSessionInMonth(ctx context.Context, month int) ([]*models.SessionCapacityResponse, error)
	UpdateSession(ctx context.Context, request []*models.SessionCapacityRequest) error
}
