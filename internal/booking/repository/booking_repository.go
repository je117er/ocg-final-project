package booking

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/booking"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type BookingRepository struct {
	conn *sql.DB
}

func NewBookingRepository(conn *sql.DB) booking.Repository {
	return &BookingRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *BookingRepository) getOne(ctx context.Context, query string, args ...interface{}) (*models.Booking, error) {
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	booking := &models.Booking{}
	err = row.Scan(&booking.ID, &booking.CustomerID, &booking.DateRegistered, &booking.DateBooked, &booking.TimePeriod, &booking.DosesCompleted,
		&booking.VaccineName, &booking.ExtendedInterval, &booking.LotNumber, &booking.ClinicName, &booking.Price, &booking.SentReminderEmail, &booking.SessionCapacityID,
		&booking.TotalBill, &booking.PaymentStatus, &booking.StockItemID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return booking, nil
}

func (repository *BookingRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Booking, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Booking, 0)
	for rows.Next() {
		booking := new(models.Booking)
		err := rows.Scan(&booking.ID, &booking.CustomerID, &booking.DateRegistered, &booking.DateBooked, &booking.TimePeriod, &booking.DosesCompleted,
			&booking.VaccineName, &booking.ExtendedInterval, &booking.LotNumber, &booking.ClinicName, &booking.Price, &booking.SentReminderEmail, &booking.SessionCapacityID,
			&booking.TotalBill, &booking.PaymentStatus, &booking.StockItemID)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, booking)
	}

	return result, nil

}

func (repository *BookingRepository) GetByCustomerID(ctx context.Context, customerID int) (*models.Booking, error) {
	query := "SELECT * FROM `booking` WHERE customer_id = ?"
	return repository.getOne(ctx, query, customerID)
}

func (repository *BookingRepository) GetByID(ctx context.Context, id int) (*models.Booking, error) {
	query := `SELECT * FROM booking WHERE id = ?`
	return repository.getOne(ctx, query, id)
}

func (repository *BookingRepository) UpdateCompleted(ctx context.Context, id int, completed int) error {
	query := `UPDATE booking SET doses_completed = ? WHERE id =?`
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return err
	}

	res, err := stmt.ExecContext(ctx, completed, id)
	if err != nil {
		logger.Error(err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
