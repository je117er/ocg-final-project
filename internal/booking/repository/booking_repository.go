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

func (repository *BookingRepository) InsertOrder(ctx context.Context, r models.OrderRequest) error {
	tx, err := repository.conn.BeginTx(ctx, nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	row := tx.QueryRow("select @customerID:=max(customer.id) + 1\nfrom customer;")
	var customerID int
	if err := row.Scan(&customerID); err != nil {
		tx.Rollback()
		return err
	}

	query := `insert into customer (email, name, address, gender, dob, phone_number, insurance_number, city, district, commune, ethnicity, nationality)
value (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err = tx.ExecContext(ctx, query,
		r.CustomerOrderRequest.Email,
		r.CustomerOrderRequest.CustomerName,
		r.CustomerOrderRequest.Address,
		r.CustomerOrderRequest.Gender,
		r.CustomerOrderRequest.DoB,
		r.CustomerOrderRequest.PhoneNumber,
		r.CustomerOrderRequest.InsuranceNumber,
		r.CustomerOrderRequest.City,
		r.CustomerOrderRequest.District,
		r.CustomerOrderRequest.Commune,
		r.CustomerOrderRequest.Ethnicity,
		r.CustomerOrderRequest.Nationality,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `INSERT INTO medical_condition (code, description, condition_status, customer_id) 
VALUES (?, ?, ?, @customerID);`
	conditionReq := r.ConditionOrderRequest
	for _, req := range conditionReq {
		_, err := tx.ExecContext(ctx, query, req.Code, req.Description, req.ConditionStatus)
		if err != nil {
			tx.Rollback()
			return nil
		}
	}

	query = `insert into booking (customer_id, date_registered, date_booked, 
                     time_period, doses_completed, vaccine_name, authorized_interval, 
                     lot_number, clinic_name, price, sent_reminder_email, session_capacity_id, 
                     total_bill, payment_status, stock_item_id)
				values (@customerID, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	_, err = tx.ExecContext(ctx, query,
		r.DateRegistered,
		r.DateBooked,
		r.TimePeriod,
		r.DosesCompleted,
		r.VaccineName,
		r.AuthorizedInterval,
		r.LotNumber,
		r.ClinicName,
		r.Price,
		r.SentReminderEmail,
		r.SessionCapacityID,
		r.TotalBill,
		0,
		r.StockItemID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `update session_capacity
set slot_left = slot_left - 1
where id = ?`
	_, err = tx.ExecContext(ctx, query, r.SessionCapacityID)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `update stock_item
set stock_left = stock_left - 1
where id = ?`
	_, err = tx.ExecContext(ctx, query, r.StockItemID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
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
