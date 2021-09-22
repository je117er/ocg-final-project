package controllers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/product"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type ProductController struct {
	ProductService product.Service
	ctx            context.Context
}

var logger = utils.SugarLog()

func NewProductController(ps product.Service, ctx context.Context, r *mux.Router) {
	controller := &ProductController{ps, ctx}
	r.Methods("GET").Path("/product/{id}").HandlerFunc(controller.GetProductByID)
	r.Methods("GET").Path("/products").HandlerFunc(controller.GetProducts)
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	results, err := c.ProductService.All(c.ctx)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		return
	}
	var responses []*models.ProductResponse
	for _, result := range results {
		responses = append(responses, result.ModelToResponse())
	}
	utils.ResponseWithJson(w, http.StatusOK, responses)
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result, err := c.ProductService.ByID(c.ctx, id)
	if err != nil {
		logger.Debug(result)
		utils.ResponseWithJson(w, http.StatusInternalServerError, utils.ResponseError{Message: "Internal Server Error"})
		return
	}

	response := result.ModelToResponse()
	utils.ResponseWithJson(w, http.StatusOK, response)
}
