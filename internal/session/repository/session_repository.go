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
