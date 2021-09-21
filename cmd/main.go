package main

import (
	"github.com/je117er/ocg-final-project/internal/server"
	"github.com/je117er/ocg-final-project/internal/utils"
	"log"
)

func main() {
	sugarLogger := utils.SugarLog()
	sugarLogger.Info("Test log info")
	sugarLogger.Debug("Test log debug")
	sugarLogger.Error("Test log err")


	log.Fatal(server.InitServer())

}
