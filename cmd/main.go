package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/ocg-final-project/internal/server"
	"github.com/je117er/ocg-final-project/internal/utils"
	"go.uber.org/zap"
	"log"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.SugarLog()
}

func main() {
	logger.Info("Start application...")

	log.Fatal(server.InitServer())
}
