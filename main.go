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
	plotterUseCase := usecase.NewPlotterUseCase(plotRepository)
	plotterHandler := handler.NewPlotterHandler(plotterUseCase)

	//Run
	router.GET("/plotter/:label/:data/:color", plotterHandler.RecieveData)
	router.Run("0.0.0.0:1988")
}
