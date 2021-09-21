package server

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/product/controllers"
	"github.com/je117er/ocg-final-project/internal/product/repository"
	"github.com/je117er/ocg-final-project/internal/product/services"
	"github.com/rs/cors"

	"log"
	"net/http"
	"time"
)

func InitServer() error {

	DB, err := sql.Open("mysql", "root:123456Aa@@tcp(localhost:6033)/vaccine-covid-19?charset=utf8mb4&parseTime=True&loc=Local")
	//DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456Aa@", "localhost", "6033", "vaccine-covid-19" ))
	if err != nil {
		log.Fatalln(err)
	}
	defer DB.Close()
	fmt.Println("you")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productRepo := repository.NewProductRepository(DB)
	serviceRepo := services.NewProductService(productRepo)

	adminRepo := admin

	r := mux.NewRouter()
	controllers.NewProductController(serviceRepo, ctx, r)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// references: https://github.com/rs/cors
	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		return err
	}
	return nil
}
