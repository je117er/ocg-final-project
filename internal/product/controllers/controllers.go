package controllers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/product"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"time"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProductController struct {
	ProductService product.Service
	ctx            context.Context
}

func NewProductController(ps product.Service, ctx context.Context, r *mux.Router) {
	controller := &ProductController{ps, ctx}
	r.Methods("GET").Path("/product/{id}").HandlerFunc(controller.GetProductByID)
	r.Methods("GET").Path("/products").HandlerFunc(controller.GetProducts)
}

type ProductResponse struct {
	ID                    string
	Name                  string
	Price                 string
	Vendor                string
	VaccineType           string
	AuthorizedAges        int32
	Dose                  string
	AntigenNature         string
	RouteOfAdministration string
	StorageRequirements   string
	AvailableFormats      string
	Diluent               string
	Adjuvant              string
	AlternateName         string
	MinimumInterval       int32
	ImmunizationSchedule  int32
	AuthorizedInterval    int32
	ExtendedInterval      int32
	Background            string
	RegulatoryActions     string
	SafetyStatus          string
	AuthorizationStatus   bool
	Trials                string
	Distribution          string
	Funding               string
	Slug                  string
	Image                 string
	LotNumber             string
	ExpiryDate            time.Time
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	results, err := c.ProductService.All(c.ctx)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, ResponseError{err.Error()})
		return
	}
	var responses []ProductResponse
	for _, result := range results {
		responses = append(responses, ProductResponse{
			ID:                    result.ID,
			Name:                  result.Name.String,
			Price:                 result.Price.String,
			Vendor:                result.Vendor.String,
			VaccineType:           result.VaccineType.String,
			AuthorizedAges:        result.AuthorizedAges.Int32,
			Dose:                  result.Dose.String,
			AntigenNature:         result.AntigenNature.String,
			RouteOfAdministration: result.RouteOfAdministration.String,
			StorageRequirements:   result.StorageRequirements.String,
			AvailableFormats:      result.AvailableFormats.String,
			Diluent:               result.Diluent.String,
			Adjuvant:              result.Adjuvant.String,
			AlternateName:         result.AlternateName.String,
			MinimumInterval:       result.MinimumInterval.Int32,
			ImmunizationSchedule:  result.ImmunizationSchedule.Int32,
			AuthorizedInterval:    result.AuthorizedInterval.Int32,
			ExtendedInterval:      result.ExtendedInterval.Int32,
			Background:            result.Background.String,
			RegulatoryActions:     result.RegulatoryActions.String,
			SafetyStatus:          result.SafetyStatus.String,
			AuthorizationStatus:   result.AuthorizationStatus.Bool,
			Trials:                result.Trials.String,
			Distribution:          result.Distribution.String,
			Funding:               result.Funding.String,
			Slug:                  result.Slug.String,
			Image:                 result.Image.String,
			LotNumber:             result.LotNumber.String,
			ExpiryDate:            result.ExpiryDate.Time,
		})
	}
	utils.ResponseWithJson(w, http.StatusOK, responses)
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	//ctx := context.Background()
	result, err := c.ProductService.ByID(c.ctx, id)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, ResponseError{"Internal Server Error"})
		return
	}

	response := &ProductResponse{
		ID:                    result.ID,
		Name:                  result.Name.String,
		Price:                 result.Price.String,
		Vendor:                result.Vendor.String,
		VaccineType:           result.VaccineType.String,
		AuthorizedAges:        result.AuthorizedAges.Int32,
		Dose:                  result.Dose.String,
		AntigenNature:         result.AntigenNature.String,
		RouteOfAdministration: result.RouteOfAdministration.String,
		StorageRequirements:   result.StorageRequirements.String,
		AvailableFormats:      result.AvailableFormats.String,
		Diluent:               result.Diluent.String,
		Adjuvant:              result.Adjuvant.String,
		AlternateName:         result.AlternateName.String,
		MinimumInterval:       result.MinimumInterval.Int32,
		ImmunizationSchedule:  result.ImmunizationSchedule.Int32,
		AuthorizedInterval:    result.AuthorizedInterval.Int32,
		ExtendedInterval:      result.ExtendedInterval.Int32,
		Background:            result.Background.String,
		RegulatoryActions:     result.RegulatoryActions.String,
		SafetyStatus:          result.SafetyStatus.String,
		AuthorizationStatus:   result.AuthorizationStatus.Bool,
		Trials:                result.Trials.String,
		Distribution:          result.Distribution.String,
		Funding:               result.Funding.String,
		Slug:                  result.Slug.String,
		Image:                 result.Image.String,
		LotNumber:             result.LotNumber.String,
		ExpiryDate:            result.ExpiryDate.Time,
	}
	utils.ResponseWithJson(w, http.StatusOK, response)
}
