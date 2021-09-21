package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	customerController "github.com/je117er/ocg-final-project/internal/customer/controller"
	customerRepo "github.com/je117er/ocg-final-project/internal/customer/repository"
	customerService "github.com/je117er/ocg-final-project/internal/customer/services"
	productController "github.com/je117er/ocg-final-project/internal/product/controllers"
	productRepo "github.com/je117er/ocg-final-project/internal/product/repository"
	productService "github.com/je117er/ocg-final-project/internal/product/services"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/rs/cors"

	"net/http"
	"time"
)

var logger = utils.SugarLog()

func InitServer() error {
	logger.Info("Start server ...")
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456Aa@", "localhost", "6033", "vaccine-covid-19"))
	if err != nil {
		logger.Fatal(err)
	}
	defer DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productRepo := productRepo.NewProductRepository(DB)
	customerRepository := customerRepo.NewCustomerRepository(DB)

	serviceRepo := productService.NewProductService(productRepo)
	customerService := customerService.NewCustomerService(customerRepository)

	r := mux.NewRouter()
	productController.NewProductController(serviceRepo, ctx, r)
	customerController.NewCustomerController(customerService, ctx, r)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// references: https://github.com/rs/cors
	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8088", handler)
	if err != nil {
		return err
	}
	return nil
}
