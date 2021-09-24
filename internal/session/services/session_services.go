package session

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/session"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type SessionService struct {
	SessionRepository session.Repository
}

func NewSessionService(sessionRepository session.Repository) session.Service {
	return &SessionService{
		SessionRepository: sessionRepository,
	}
}

var logger = utils.SugarLog()

func (service *SessionService) GetAllSessionInMonth(ctx context.Context, month int) ([]*models.SessionCapacityResponse, error) {
	reusult, err := service.SessionRepository.GetAllSessionInMonth(ctx, month)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	response := make([]*models.SessionCapacityResponse, 0)
	for _, v := range reusult {
		response = append(response, v.Entity2Response())
	}
	return response, nil
}

func (service *SessionService) UpdateSession(ctx context.Context, requests []*models.SessionCapacityRequest) error {
	for _, v := range requests {
		service.SessionRepository.Update(ctx, v)
	}
	return nil
}
