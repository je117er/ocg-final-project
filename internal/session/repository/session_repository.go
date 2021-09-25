package session

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/session"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type SessionRepository struct {
	conn *sql.DB
}

func NewSessionRepository(conn *sql.DB) session.Repository {
	return &SessionRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *SessionRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.SessionCapacity, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.SessionCapacity, 0)
	for rows.Next() {
		session := new(models.SessionCapacity)
		err := rows.Scan(&session.ID, &session.ClinicID, &session.Capacity, &session.Type,
			&session.Status, &session.CurrentDate, &session.SlotLeft)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, session)
	}

	return result, nil

}

func (repository *SessionRepository) GetSessionByClinic(ctx context.Context, clinicName string) ([]*models.SessionByClinic, error) {
	query := `select c.id, s.id as session_id, s.type as time_period, s.current_date from session_capacity as s
		join clinic c on s.clinic_id = c.id
		where s.slot_left > 0 && s.status <> 0 && c.name like ?`
	logger.Info(query)
	rows, err := repository.conn.QueryContext(ctx, query, "%" + clinicName + "%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*models.SessionByClinic, 0)
	for rows.Next() {
		result := new(models.SessionByClinic)
		if err := rows.Scan(&result.ClinicID, &result.SessionID, &result.TimePeriod, &result.CurrentDate); err != nil {
			logger.Error(err)
			continue
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}
	return results, nil
}
func (repository *SessionRepository) GetAllSessionInMonth(ctx context.Context, month int) ([]*models.SessionCapacity, error) {
	query := "SELECT * FROM session_capacity WHERE MONTH(`current_date`) = ?"
	return repository.fetch(ctx, query, month)
}

func (repository *SessionRepository) Update(ctx context.Context, request *models.SessionCapacityRequest) error {
	query := `UPDATE session_capacity SET capacity = ?, status = ? WHERE id = ?`
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, request.Capacity, request.Status, request.ID)
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
