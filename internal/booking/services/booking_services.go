package booking

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/booking"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/stock"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type BookingService struct {
	BookingRepository booking.Repository
	SockRepository    stock.Repository
}

func NewBookingService(bookingRepository booking.Repository, sockRepository stock.Repository) booking.Service {
	return &BookingService{bookingRepository, sockRepository}
}

var logger = utils.SugarLog()

func (service *BookingService) GetVaccinationByCustomerID(ctx context.Context, customerID int) ([]models.VaccinationResponse, error) {
	res, err := service.BookingRepository.GetByCustomerID(ctx, customerID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	stock, err := service.SockRepository.GetByID(ctx, res.StockItemID.String)

	immunizationSchedule := stock.ImmunizationSchedule.Int64
	vaccinations := make([]models.VaccinationResponse, 0)

	currentTime := res.DateBooked.Time
	for i := 1; i <= int(immunizationSchedule); i++ {
		currentTime = currentTime.AddDate(0, 0, (i-1)*int(res.ExtendedInterval.Int64))

		vaccinationResponse := models.VaccinationResponse{
			BookingID:        int(res.CustomerID),
			VaccineName:      res.VaccineName.String,
			Dose:             i,
			RegistrationTime: res.DateRegistered.Time,
			Completed:        int(res.DosesCompleted.Int64) >= i,
			VaccinationTime:  currentTime,
			Clinic:           res.ClinicName.String,
		}

		vaccinations = append(vaccinations, vaccinationResponse)
	}

	return vaccinations, nil
}

func (service *BookingService) UpdateCustomer(ctx context.Context, request models.VaccinationRequest) ([]models.VaccinationResponse, error) {
	bookingModel, err := service.BookingRepository.GetByID(ctx, request.BookingID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = service.BookingRepository.UpdateCompleted(ctx, request.BookingID, request.Completed)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return service.GetVaccinationByCustomerID(ctx, int(bookingModel.CustomerID))
}
