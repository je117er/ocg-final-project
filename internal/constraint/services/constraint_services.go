package constraint

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/constraint"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type ConstraintService struct {
	constraintRepository constraint.Repository
}

func NewConstraintService(constraintRepository constraint.Repository) constraint.Service {
	return &ConstraintService{
		constraintRepository: constraintRepository,
	}
}

var logger = utils.SugarLog()

func (service *ConstraintService) GetAll(ctx context.Context) ([]*models.ConstraintResponse, error) {
	res, err := service.constraintRepository.Fetch(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	result := make([]*models.ConstraintResponse, 0)
	for _, v := range res {
		result = append(result, v.Entity2Response())
	}
	return result, nil
}
