package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/je117er/ocg-final-project/internal/models"
)

type ProductRepository struct {
	Conn *sql.DB
}

func NewProductRepository(Conn *sql.DB) *ProductRepository {
	return &ProductRepository{Conn}
}

func (p *ProductRepository) ByID(ctx context.Context, id string) (models.Product, error) {
	var product models.Product
 	//query := `# select * from product where id = uuid_to_bin(?)`
	//query := `SELECT BIN_TO_UUID(id) id, name, price, vendor, vaccine_type, authorized_ages, dose, antigen_nature, route_of_administration, storage_requirements, available_formats, diluent, adjuvant, alternate_name, minimum_interval, immunization_schedule, authorized_interval, extended_interval, background, regulatory_actions, safety_status, authorization_status, trials, distribution, funding, slug, image, lot_number, expiry_date FROM product WHERE bin_to_uuid(id) = ?;`
	query := `select * from product where bin_to_uuid(id) = ?`
	err := p.Conn.QueryRowContext(ctx, query, id).Scan(
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
	)
	switch {
	case err == sql.ErrNoRows:
		return models.Product{}, fmt.Errorf("No product with id %s\n", id)
	case err != nil:
		return models.Product{}, fmt.Errorf("Query error: %v\n", err)
	}
	return product, nil
}

func (p *ProductRepository) All(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	//query := `SELECT BIN_TO_UUID(id) id, name, price, vendor, vaccine_type, authorized_ages, dose, antigen_nature, route_of_administration, storage_requirements, available_formats, diluent, adjuvant, alternate_name, minimum_interval, immunization_schedule, authorized_interval, extended_interval, background, regulatory_actions, safety_status, authorization_status, trials, distribution, funding, slug, image, lot_number, expiry_date FROM product`
	query := `select * from product where bin_to_uuid(id) = ?`
	fmt.Println("you are hereee")
	rows, err := p.Conn.QueryContext(ctx, query)
	fmt.Println("you are hereee")
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
			return nil, err
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
