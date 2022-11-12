package main

import (
	"main/src/infrastructures/persistence"
	"main/src/presentation/handler"
	"main/src/usecase"
)

func main() {
	plotRepository := persistence.NewPlotPersistence()
	httpRequestRepository := persistence.NewHttpRequestPersistence()
	cloudWatchUseCase := usecase.NewCloudWatchUseCase(plotRepository, httpRequestRepository)
	cloudWatchHandler := handler.NewCloudWatchHandler(cloudWatchUseCase)
	cloudWatchHandler.RunCheckSpeedInsight()
}
