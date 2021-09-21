package repository

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
)

type AdminRepository struct {
	Conn *sql.DB
}

func NewAdminRepository(Conn *sql.DB) *AdminRepository {
	return &AdminRepository{Conn}
}

func (a *AdminRepository) Fetch(ctx context.Context, query string, args... interface{}) (*models.Admin, error) {
	stmt, err := a.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	admin := &models.Admin{}
	if err := row.Scan(&admin.ID, &admin.Username, &admin.HashedPassword); err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *AdminRepository) ByUsername(ctx context.Context, username string) (*models.Admin, error) {
	query := `SELECT * FROM admin WHERE username = ?`
	res, err := a.Fetch(ctx, query, username)
	if err != nil {
		return nil, err
	}
	return res, nil
}
