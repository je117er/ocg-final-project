package clinic

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/clinic"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type ClinicService struct {
	clinicRepository clinic.Repository
}

func NewClinicService(clinicRepository clinic.Repository) clinic.Service {
	return &ClinicService{
		clinicRepository: clinicRepository,
	}
}

var logger = utils.SugarLog()

func (service *ClinicService) GetAll(ctx context.Context) ([]*models.ClinicResponse, error) {
	res, err := service.clinicRepository.Fetch(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	result := make([]*models.ClinicResponse, 0)
	for _, v := range res {
		result = append(result, v.Entity2Response())
	}
	return result, nil
}

func (service *ClinicService) FindByID(ctx context.Context, id string) (*models.ClinicResponse, error) {
	res, err := service.clinicRepository.GetByID(ctx, id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res.Entity2Response(), err
}

func (service *ClinicService) UpdateClinic(ctx context.Context, customerRequest models.ClinicRequest) (*models.ClinicResponse, error) {
	_, err := service.clinicRepository.GetByID(ctx, customerRequest.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	clinicModel, err := service.clinicRepository.Update(ctx, customerRequest)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return clinicModel.Entity2Response(), nil
}
