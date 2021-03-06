package repository

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

var logger = utils.SugarLog()

type ProductRepository struct {
	Conn *sql.DB
}

func NewProductRepository(Conn *sql.DB) *ProductRepository {
	return &ProductRepository{Conn}
}

func (p *ProductRepository) Fetch(ctx context.Context, query string, args ...interface{}) (*models.Product, error) {
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	product := &models.Product{}
	if err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Vendor,
		&product.VaccineType,
		&product.AuthorizedAges,
		&product.Dose,
		&product.AntigenNature,
		&product.RouteOfAdministration,
		&product.StorageRequirements,
		&product.AvailableFormats,
		&product.Diluent,
		&product.Adjuvant,
		&product.AlternateName,
		&product.MinimumInterval,
		&product.ImmunizationSchedule,
		&product.AuthorizedInterval,
		&product.ExtendedInterval,
		&product.Background,
		&product.RegulatoryActions,
		&product.SafetyStatus,
		&product.AuthorizationStatus,
		&product.Trials,
		&product.Distribution,
		&product.Funding,
		&product.Slug,
		&product.Image,
		&product.LotNumber,
		&product.ExpiryDate,
	); err != nil {
		return nil, err
	}
	return product, nil
}
func (p *ProductRepository) ByID(ctx context.Context, id string) (*models.Product, error) {
	query := `SELECT * FROM product where id = ?`
	return p.Fetch(ctx, query, id)
}

func (p *ProductRepository) AllNameID(ctx context.Context) ([]models.ProductNameID, error) {
	query := `select id, name, immunization_schedule, price from product`
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductNameID
	var product models.ProductNameID

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.ImmunizationSchedule, &product.Price); err != nil {
			logger.Error(err)
			continue
		}
		products = append(products, product)
	}
	rerr := rows.Close()
	if rerr != nil {
		logger.Error(err)
		return nil, rerr
	}
	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) All(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	//query := `SELECT BIN_TO_UUID(id) id, name, price, vendor, vaccine_type, authorized_ages, dose, antigen_nature, route_of_administration, storage_requirements, available_formats, diluent, adjuvant, alternate_name, minimum_interval, immunization_schedule, authorized_interval, extended_interval, background, regulatory_actions, safety_status, authorization_status, trials, distribution, funding, slug, image, lot_number, expiry_date FROM product`
	query := `select * from product`
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()
	var product models.Product
	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Vendor,
			&product.VaccineType,
			&product.AuthorizedAges,
			&product.Dose,
			&product.AntigenNature,
			&product.RouteOfAdministration,
			&product.StorageRequirements,
			&product.AvailableFormats,
			&product.Diluent,
			&product.Adjuvant,
			&product.AlternateName,
			&product.MinimumInterval,
			&product.ImmunizationSchedule,
			&product.AuthorizedInterval,
			&product.ExtendedInterval,
			&product.Background,
			&product.RegulatoryActions,
			&product.SafetyStatus,
			&product.AuthorizationStatus,
			&product.Trials,
			&product.Distribution,
			&product.Funding,
			&product.Slug,
			&product.Image,
			&product.LotNumber,
			&product.ExpiryDate,
		); err != nil {
			continue
		}
		products = append(products, product)
	}
	rerr := rows.Close()
	if rerr != nil {
		return nil, rerr
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil

}
