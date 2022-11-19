package main

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/src/infrastructures/persistence"
	"main/src/presentation/handler"
	"main/src/usecase"
)

func main() {

	router := gin.Default()
	db := config.Connect()
	defer db.Close()

	//DI
	plotRepository := persistence.NewPlotPersistence(db)
	cloudWatchUseCase := usecase.NewCloudWatchUseCase(plotRepository)
	cloudWatchHandler := handler.NewCloudWatchHandler(cloudWatchUseCase)

	//Run
	router.GET("/plotter/:label/:data/:color", cloudWatchHandler.PlotData)
	router.Run("localhost:8080")
}
