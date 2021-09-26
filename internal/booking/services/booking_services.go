package booking

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/booking"
	"github.com/je117er/ocg-final-project/internal/constraint"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/session"
	"github.com/je117er/ocg-final-project/internal/stock"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type BookingService struct {
	BookingRepository    booking.Repository
	SockRepository       stock.Repository
	CustomerRepository   customer.Repository
	ConstraintRepository constraint.Repository
	SessionRepository    session.Repository
}

func NewBookingService(bookingRepository booking.Repository, sockRepository stock.Repository, customerRepository customer.Repository, constraintRepository constraint.Repository, sessionRepository session.Repository) booking.Service {
	return &BookingService{
		BookingRepository:    bookingRepository,
		SockRepository:       sockRepository,
		CustomerRepository:   customerRepository,
		ConstraintRepository: constraintRepository,
		SessionRepository:    sessionRepository,
	}
}

var logger = utils.SugarLog()

func (service *BookingService) InsertOrder(ctx context.Context, r models.OrderRequest) error {
	id, txCustomer, err := service.CustomerRepository.SaveCustomer(ctx, r.CustomerOrderRequest)
	if err != nil {
		logger.Error(err)
		return err
	}

	txConstraint, err := service.ConstraintRepository.SaveConstraint(ctx, r.ConditionOrderRequest, id)
	if err != nil {
		logger.Error(err)
		service.rollback(ctx, id)
		return err
	}

	txSession, err := service.SessionRepository.UpdateCapacity(ctx, 1, r.SessionCapacityID)
	if err != nil {
		service.rollback(ctx, id, txConstraint)
		logger.Error(err)
		return err
	}

	txStock, err := service.SockRepository.UpdateStock(ctx, 1, r.StockItemID)
	if err != nil {
		service.rollback(ctx, id, txConstraint, txSession)
		logger.Error(err)
		return err
	}

	txBooking, err := service.BookingRepository.SaveOrder(ctx, r, id)
	if err != nil {
		service.rollback(ctx, id, txConstraint, txSession, txStock)
		logger.Error(err)
		return err
	}

	return func(txList ...*sql.Tx) error {
		for _, tx := range txList {
			if err := tx.Commit(); err != nil {
				logger.Error(err)
				service.rollback(ctx, id, txConstraint, txBooking, txSession)
				return err
			}
		}
		return nil
	}(txCustomer, txConstraint, txBooking, txSession, txStock)
}

func (service *BookingService) rollback(ctx context.Context, customerID int, txList ...*sql.Tx) {
	service.CustomerRepository.DeleteByID(ctx, customerID)
	for _, tx := range txList {
		tx.Rollback()
	}
}

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
		currentTime = currentTime.AddDate(0, 0, (i-1)*int(res.AuthorizedInterval.Int64))

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
