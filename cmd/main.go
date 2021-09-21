package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/product/controllers"
	"github.com/je117er/ocg-final-project/internal/product/repository"
	"github.com/je117er/ocg-final-project/internal/product/services"
	"github.com/je117er/ocg-final-project/internal/utils"
	"log"
	"net/http"
)

func main() {
	sugarLogger := utils.SugarLog()
	sugarLogger.Info("Test log info")
	sugarLogger.Debug("Test log debug")
	sugarLogger.Error("Test log err")

	DB, err := sql.Open("mysql", "root:123456Aa@@tcp(localhost:6033)/vaccine-covid-19?charset=utf8mb4&parseTime=True&loc=Local")
	//DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456Aa@", "localhost", "6033", "vaccine-covid-19" ))
	if err != nil {
		log.Fatalln(err)
	}
	defer DB.Close()
	fmt.Println("you")
	sugarLogger.Error("Test log err")

	productRepo := repository.NewProductRepository(DB)
	serviceRepo := services.NewProductService(productRepo)

	r := mux.NewRouter()
	controllers.NewProductController(serviceRepo, r)

	sugarLogger.Fatal(http.ListenAndServe(":8000", r))
}
